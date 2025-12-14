package main

import (
	"log"

	"workflow_engine/internal/config"
	"workflow_engine/internal/db"
	"workflow_engine/internal/engine"
	"workflow_engine/internal/queue"
)

func main() {
	cfg := config.Load()

	database, err := db.Connect(cfg.DBUrl)
	if err != nil {
		log.Fatal("worker db connection failed:", err)
	}

	q := queue.NewRedisQueue(cfg.RedisUrl)

	log.Println("Worker started...")
	engine.StartWorker(database, q)
}
