package db

import (
    "io/ioutil"
    "log"
    "strings"

    "github.com/jmoiron/sqlx"
)

func RunMigrations(db *sqlx.DB, path string) {
    data, err := ioutil.ReadFile(path)
    if err != nil {
        log.Fatalf("could not read migrations: %v", err)
    }

    queries := strings.Split(string(data), ";")

    for _, q := range queries {
        q = strings.TrimSpace(q)
        if q == "" {
            continue
        }
        _, err := db.Exec(q)
        if err != nil {
            log.Fatalf("migration failed: %v", err)
        }
    }

    log.Println("migrations applied successfully")
}
