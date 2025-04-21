package config

import (
    "database/sql"
    "fmt"
    "log"   
    "os"
    _ "github.com/go-sql-driver/mysql"
)
    
var DB *sql.DB

func Connect() {
    dbHost := os.Getenv("DB_HOST")
    dbPort := os.Getenv("DB_PORT")
    dbUser := os.Getenv("DB_USER")
    dbPassword := os.Getenv("DB_PASSWORD")
    dbName := os.Getenv("DB_NAME")

    dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
        dbUser, dbPassword, dbHost, dbPort, dbName)

    var err error
    DB, err = sql.Open("mysql", dsn)
    if err != nil {
        log.Fatalf("❌ Failed to connect to the database: %v", err)
    }

    err = DB.Ping()
    if err != nil {
        log.Fatalf("❌ Database ping failed: %v", err)
    }

    log.Println("✅ Connected to the database successfully!")
}