package models

type TriggerEvent struct {
    ID     int         `db:"id"`
    Payload map[string]interface{} `db:"payload"`
}
