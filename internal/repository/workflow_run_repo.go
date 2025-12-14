package repository

import (
    "time"
    "github.com/jmoiron/sqlx"
)

type WorkflowRunRepo struct {
    db *sqlx.DB
}

func NewWorkflowRunRepo(db *sqlx.DB) *WorkflowRunRepo {
    return &WorkflowRunRepo{db: db}
}

func (r *WorkflowRunRepo) CreateWorkflowRun(workflowID, eventID int) (int, error) {
    var id int
    query := `
        INSERT INTO workflow_runs (workflow_id, trigger_event_id, status, started_at)
        VALUES ($1, $2, 'running', $3)
        RETURNING id;
    `
    err := r.db.QueryRow(query, workflowID, eventID, time.Now()).Scan(&id)
    return id, err
}
func (r *WorkflowRunRepo) MarkSuccess(runID int) error {
    query := `
        UPDATE workflow_runs
        SET status='success', finished_at=NOW()
        WHERE id=$1;
    `
    _, err := r.db.Exec(query, runID)
    return err
}

func (r *WorkflowRunRepo) MarkFailed(runID int, errMsg string) error {
    query := `
        UPDATE workflow_runs
        SET status='failed', error_message=$1, finished_at=NOW()
        WHERE id=$2;
    `
    _, err := r.db.Exec(query, errMsg, runID)
    return err
}
