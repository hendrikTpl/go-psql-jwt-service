package database

import (
    "database/sql"
    "fmt"
    "backend/config"

    _ "github.com/lib/pq" // PostgreSQL driver
)

func ConnectDB(cfg *config.Config) (*sql.DB, error) {
    dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable",
        cfg.DBHost, cfg.DBUser, cfg.DBPassword, cfg.DBName)
    db, err := sql.Open("postgres", dsn)
    if err != nil {
        return nil, err
    }

    // Test the connection
    if err := db.Ping(); err != nil {
        return nil, err
    }

    return db, nil
}
