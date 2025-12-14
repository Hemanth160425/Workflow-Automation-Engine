package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

type WorkflowStep struct {
	ID           int             `db:"id" json:"id"`
	WorkflowID   int             `db:"workflow_id" json:"workflow_id"`
	StepNumber   int             `db:"step_number" json:"step_number"`
	ActionType   string          `db:"action_type" json:"action_type"`
	ActionConfig json.RawMessage `db:"action_config" json:"action_config"`
	CreatedAt    string          `db:"created_at" json:"created_at"`
}

// Value implements the driver.Valuer interface for WorkflowStep
type ActionConfig map[string]interface{}

// Value implements the driver.Valuer interface
func (a ActionConfig) Value() (driver.Value, error) {
	return json.Marshal(a)
}

// Scan implements the sql.Scanner interface
func (a *ActionConfig) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(b, &a)
}
