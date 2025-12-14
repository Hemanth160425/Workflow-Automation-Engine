package engine

import (
	"encoding/json"
	"fmt"
	"log"

	"workflow_engine/internal/models"
	"workflow_engine/internal/queue"
	"workflow_engine/internal/repository"

	"github.com/jmoiron/sqlx"
)

const JobQueue = "workflow_jobs"
const MaxRetries = 3

func StartWorker(db *sqlx.DB, q *queue.RedisQueue) {
	for {
		jobJSON, err := q.Pop(JobQueue)
		if err != nil {
			continue
		}

		var job models.Job
		if err := json.Unmarshal([]byte(jobJSON), &job); err != nil {
			log.Println("invalid job:", err)
			continue
		}

		ExecuteWorkflow(db, q, job)
	}
}

func ExecuteWorkflow(db *sqlx.DB, q *queue.RedisQueue, job models.Job) {

	runRepo := repository.NewWorkflowRunRepo(db)
	eventRepo := repository.NewTriggerEventRepo(db)
	stepRepo := repository.NewWorkflowStepRepo(db)
	stepRunRepo := repository.NewStepRunRepo(db)

	// 1. Create workflow_run
	runID, err := runRepo.CreateWorkflowRun(job.WorkflowID, job.EventID)
	if err != nil {
		log.Println("failed to create workflow_run:", err)
		return
	}

	// 2. Load trigger payload (initial input)
	event, err := eventRepo.GetEvent(job.EventID)
	if err != nil {
		_ = runRepo.MarkFailed(runID, err.Error())
		return
	}

	currentInput := event.Payload

	// 3. Load workflow steps
	steps, err := stepRepo.GetWorkflowSteps(job.WorkflowID)
	if err != nil {
		_ = runRepo.MarkFailed(runID, err.Error())
		return
	}

	// 4. Execute steps in order
	for _, step := range steps {

		retries := 0

		for {
			// 4.1 Create step_run BEFORE execution
			stepRunID, err := stepRunRepo.CreateStepRun(
				runID,
				step.StepNumber,
				currentInput,
			)
			if err != nil {
				_ = runRepo.MarkFailed(runID, err.Error())
				return
			}

			// 4.2 Unmarshal action config
			var configMap map[string]interface{}
			if err := json.Unmarshal(step.ActionConfig, &configMap); err != nil {
				_ = runRepo.MarkFailed(runID, fmt.Sprintf("failed to unmarshal action config: %v", err))
				return
			}

			// 4.3 Execute action
			output, err := ExecuteAction(
				step.ActionType,
				configMap,
				currentInput,
			)

			if err != nil {
				_ = stepRunRepo.FailStepRun(stepRunID, err.Error())
				retries++

				if retries >= MaxRetries {
					_ = runRepo.MarkFailed(runID, err.Error())
					return
				}

				// retry same step
				continue
			}

			// 4.3 Mark step success
			_ = stepRunRepo.CompleteStepRun(stepRunID, output)

			// 4.4 Chain output → next step input
			currentInput = output
			break
		}
	}

	// 5. Mark workflow success
	_ = runRepo.MarkSuccess(runID)
}
