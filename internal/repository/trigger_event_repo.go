package repository

import (
    "workflow_engine/internal/models"
    "github.com/jmoiron/sqlx"
)

type TriggerEventRepo struct {
    db *sqlx.DB
}

func NewTriggerEventRepo(db *sqlx.DB) *TriggerEventRepo {
    return &TriggerEventRepo{db: db}
}

func (r *TriggerEventRepo) GetEvent(id int) (*models.TriggerEvent, error) {
    var event models.TriggerEvent
    query := `SELECT id, payload FROM trigger_events WHERE id=$1`
    err := r.db.Get(&event, query, id)
    return &event, err
}
