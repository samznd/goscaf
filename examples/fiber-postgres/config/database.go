package config

import (
    "database/sql"
    "fmt"
    "log"
    "os"
    _ "github.com/lib/pq"
)

var DB *sql.DB

func Connect() {
    dbHost := os.Getenv("DB_HOST")
    dbPort := os.Getenv("DB_PORT")
    dbUser := os.Getenv("DB_USER")
    dbPassword := os.Getenv("DB_PASSWORD")
    dbName := os.Getenv("DB_NAME")

    dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
        dbHost, dbPort, dbUser, dbPassword, dbName)

    var err error
    DB, err = sql.Open("postgres", dsn)
    if err != nil {
        log.Fatalf("❌ Failed to connect to the database: %v", err)
    }

    err = DB.Ping()
    if err != nil {
        log.Fatalf("❌ Database ping failed: %v", err)
    }

    log.Println("✅ Connected to the database successfully!")
}