package models

type StepRun struct {
    ID            int                    `db:"id"`
    WorkflowRunID int                    `db:"workflow_run_id"`
    StepNumber    int                    `db:"step_number"`
    Status        string                 `db:"status"`
    InputData     map[string]interface{} `db:"input_data"`
    OutputData    map[string]interface{} `db:"output_data"`
    ErrorMessage  *string                `db:"error_message"`
}
