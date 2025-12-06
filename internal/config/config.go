package config

import (
    "os"

    "github.com/joho/godotenv"
)

type Config struct {
    DBUrl    string
    RedisUrl string
    Port     string
}

func Load() *Config {
    _ = godotenv.Load() // loads .env file if exists

    return &Config{
        DBUrl:    os.Getenv("DATABASE_URL"),
        RedisUrl: os.Getenv("REDIS_URL"),
        Port:     os.Getenv("PORT"),
    }
}
