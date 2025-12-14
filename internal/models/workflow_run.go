package models

type WorkflowRun struct {
    ID             int     `db:"id"`
    WorkflowID     int     `db:"workflow_id"`
    TriggerEventID int     `db:"trigger_event_id"`
    Status         string  `db:"status"`
    ErrorMessage   *string `db:"error_message"`
    StartedAt      *string `db:"started_at"`
    FinishedAt     *string `db:"finished_at"`
}
