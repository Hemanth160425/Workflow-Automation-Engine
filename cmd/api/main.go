package main

import (
    "log"

    "workflow_engine/internal/config"
    "workflow_engine/internal/db"
)

func main() {
    cfg := config.Load()

    database, err := db.Connect(cfg.DBUrl)
    if err != nil {
        log.Fatal("db connection failed: ", err)
    }

    db.RunMigrations(database, "migrations/001_create_tables.sql")

    log.Println("API server starting...")
}
