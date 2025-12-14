package repository

import (
    "github.com/jmoiron/sqlx"
)

type StepRunRepo struct {
    db *sqlx.DB
}

func NewStepRunRepo(db *sqlx.DB) *StepRunRepo {
    return &StepRunRepo{db: db}
}

func (r *StepRunRepo) CreateStepRun(runID int, stepNumber int, input map[string]interface{}) (int, error) {
    var id int
    query := `
        INSERT INTO step_runs (workflow_run_id, step_number, status, input_data, started_at)
        VALUES ($1, $2, 'running', $3, NOW())
        RETURNING id;
    `
    err := r.db.QueryRow(query, runID, stepNumber, input).Scan(&id)
    return id, err
}

func (r *StepRunRepo) CompleteStepRun(id int, output map[string]interface{}) error {
    query := `
        UPDATE step_runs
        SET status='success', output_data=$1, finished_at=NOW()
        WHERE id=$2;
    `
    _, err := r.db.Exec(query, output, id)
    return err
}

func (r *StepRunRepo) FailStepRun(id int, errMsg string) error {
    query := `
        UPDATE step_runs
        SET status='failed', error_message=$1, finished_at=NOW()
        WHERE id=$2;
    `
    _, err := r.db.Exec(query, errMsg, id)
    return err
}
