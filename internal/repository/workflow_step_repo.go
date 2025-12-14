package repository

import (
	"workflow_engine/internal/models"

	"github.com/jmoiron/sqlx"
)

type WorkflowStepRepo struct {
	db *sqlx.DB
}

func NewWorkflowStepRepo(db *sqlx.DB) *WorkflowStepRepo {
	return &WorkflowStepRepo{db: db}
}

func (r *WorkflowStepRepo) GetWorkflowSteps(workflowID int) ([]models.WorkflowStep, error) {
	var steps []models.WorkflowStep
	query := `
		SELECT id, workflow_id, step_number, action_type, action_config, created_at
		FROM workflow_steps
		WHERE workflow_id = $1
		ORDER BY step_number ASC
	`
	err := r.db.Select(&steps, query, workflowID)
	return steps, err
}
