package models

type Job struct {
	WorkflowID int `json:"workflow_id"`
	EventID    int `json:"event_id"`
}
