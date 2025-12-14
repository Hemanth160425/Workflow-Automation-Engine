package main

import (
	"log"

	"workflow_engine/internal/config"
	"workflow_engine/internal/db"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.Load()

	database, err := db.Connect(cfg.DBUrl)
	if err != nil {
		log.Fatal("db connection failed:", err)
	}

	db.RunMigrations(database, "migrations/001_create_tables.sql")

	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	log.Println("API running on port", cfg.Port)
	r.Run(":" + cfg.Port)
}
