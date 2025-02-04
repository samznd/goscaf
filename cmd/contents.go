package cmd

import "fmt"

func getMainFile(backend string, projectName string) string {
	if backend == "Fiber" {
		return fmt.Sprintf(`package main

		import (
			"fmt"
			"log"
			"github.com/gofiber/fiber/v2"
			"%s/config/database"
		)
		
		func main() {
			database.Connect()
			app := fiber.New()
		
			// Define routes
			app.Get("/ping", func(c *fiber.Ctx) error {
				return c.JSON(fiber.Map{"message": "pong"})
			})
		
			log.Println("🚀 Fiber server is running on http://localhost:3000")
			app.Listen(":3000")
		}`, projectName)
	}
	return fmt.Sprintf(`package main

import (
    "fmt"
    "log"
    "net/http"
    "github.com/gin-gonic/gin"
	"%s/config/database"
)

func main() {
	database.Connect()
	
    r := gin.Default()

    // Define routes
    r.GET("/ping", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{"message": "pong"})
    })

    log.Println("🚀 Gin server is running on http://localhost:3000")
    r.Run(":3000")
}`, projectName)
}

func getDatabaseFile(database string, orm string) string {
	if orm != "None" {
		return SetupORM(orm, database)
	}
	if database == "Postgres" {
		return `package database

import (
    "database/sql"
    "fmt"
    "log"
    "os"

    _ "github.com/lib/pq"
)

var DB *sql.DB

func Connect() {
    dbHost := getEnv("DB_HOST", "localhost")
    dbPort := getEnv("DB_PORT", "5432")
    dbUser := getEnv("DB_USER", "postgres")
    dbPassword := getEnv("DB_PASSWORD", "password")
    dbName := getEnv("DB_NAME", "mydb")

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

func getEnv(key, defaultValue string) string {
    value, exists := os.LookupEnv(key)
    if !exists {
        return defaultValue
    }
    return value
}`
	}
	if database == "MySQL" {
		return `package database

import (
    "database/sql"
    "fmt"
    "log"   
    "os"
    "github.com/go-sql-driver/mysql"
)
    
var DB *sql.DB

func Connect() {
    dbHost := getEnv("DB_HOST", "localhost")
    dbPort := getEnv("DB_PORT", "3306")
    dbUser := getEnv("DB_USER", "root")
    dbPassword := getEnv("DB_PASSWORD", "password")
    dbName := getEnv("DB_NAME", "mydb")

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
}`
	}

	if database == "SQLite" {
		return `package database

import (
    "database/sql"
    "log"
    "os"
    _ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB


func Connect() {
    dbName := getEnv("DB_NAME", "mydb")

    var err error
    DB, err = sql.Open("sqlite3", dbName)
    if err != nil {
        log.Fatalf("❌ Failed to connect to the database: %v", err)
    }

    err = DB.Ping()
    if err != nil {
        log.Fatalf("❌ Database ping failed: %v", err)
    }

    log.Println("✅ Connected to the database successfully!")
}

func getEnv(key, defaultValue string) string {
    value, exists := os.LookupEnv(key)
    if !exists {
        return defaultValue
    }
    return value
}`
	}
	return "None"
}

func SetupORM(orm string, database string) string {
	if orm == "GORM" {
		if database == "Postgres" {
			return `package database

import (
    "fmt"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
    "log"
    "os"
)

var DB *gorm.DB

func Connect() {
    dbHost := getEnv("DB_HOST", "localhost")
    dbPort := getEnv("DB_PORT", "5432")
    dbUser := getEnv("DB_USER", "postgres")
    dbPassword := getEnv("DB_PASSWORD", "password")
    dbName := getEnv("DB_NAME", "mydb")

    dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
        dbHost, dbPort, dbUser, dbPassword, dbName)

    var err error
    DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatalf("❌ Failed to connect to the database: %v", err)
    }

    log.Println("✅ Connected to the PostgreSQL database successfully!")
}`
		}
		if database == "MySQL" {
			return `package database

import (
    "fmt"
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
    "log"
    "os"
)

var DB *gorm.DB

func Connect() {
    dbHost := getEnv("DB_HOST", "localhost")
    dbPort := getEnv("DB_PORT", "3306")
    dbUser := getEnv("DB_USER", "root")
    dbPassword := getEnv("DB_PASSWORD", "password")
    dbName := getEnv("DB_NAME", "mydb")

    dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
        dbUser, dbPassword, dbHost, dbPort, dbName)

    var err error
    DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatalf("❌ Failed to connect to the database: %v", err)
    }

    log.Println("✅ Connected to the MySQL database successfully!")
}`
		}
		if database == "SQLite" {
			return `package database

import (
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
    "log"
    "os"
)

var DB *gorm.DB

func Connect() {
    dbName := getEnv("DB_NAME", "mydb.db")

    var err error
    DB, err = gorm.Open(sqlite.Open(dbName), &gorm.Config{})
    if err != nil {
        log.Fatalf("❌ Failed to connect to the database: %v", err)
    }

    log.Println("✅ Connected to the SQLite database successfully!")
}`
		}
	} else if orm == "XORM" {
		if database == "Postgres" {
			return `package database

import (
    "fmt"
    "log"
    "os"
    "xorm.io/xorm"
    _ "github.com/lib/pq"
)

var DB *xorm.Engine

func Connect() {
    dbHost := getEnv("DB_HOST", "localhost")
    dbPort := getEnv("DB_PORT", "5432")
    dbUser := getEnv("DB_USER", "postgres")
    dbPassword := getEnv("DB_PASSWORD", "password")
    dbName := getEnv("DB_NAME", "mydb")

    dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
        dbHost, dbPort, dbUser, dbPassword, dbName)

    var err error
    DB, err = xorm.NewEngine("postgres", dsn)
    if err != nil {
        log.Fatalf("❌ Failed to connect to the database: %v", err)
    }

    if err := DB.Ping(); err != nil {
        log.Fatalf("❌ Database ping failed: %v", err)
    }

    log.Println("✅ Connected to the PostgreSQL database successfully!")
}`
		}
		if database == "MySQL" {
			return `package database

import (
    "fmt"
    "log"
    "os"
    "xorm.io/xorm"
    _ "github.com/go-sql-driver/mysql"
)

var DB *xorm.Engine

func Connect() {
    dbHost := getEnv("DB_HOST", "localhost")
    dbPort := getEnv("DB_PORT", "3306")
    dbUser := getEnv("DB_USER", "root")
    dbPassword := getEnv("DB_PASSWORD", "password")
    dbName := getEnv("DB_NAME", "mydb")

    dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4",
        dbUser, dbPassword, dbHost, dbPort, dbName)

    var err error
    DB, err = xorm.NewEngine("mysql", dsn)
    if err != nil {
        log.Fatalf("❌ Failed to connect to the database: %v", err)
    }

    if err := DB.Ping(); err != nil {
        log.Fatalf("❌ Database ping failed: %v", err)
    }

    log.Println("✅ Connected to the MySQL database successfully!")
}`
		}
		if database == "SQLite" {
			return `package database

import (
    "log"
    "os"
    "xorm.io/xorm"
    _ "github.com/mattn/go-sqlite3"
)

var DB *xorm.Engine

func Connect() {
    dbName := getEnv("DB_NAME", "mydb.db")

    var err error
    DB, err = xorm.NewEngine("sqlite3", dbName)
    if err != nil {
        log.Fatalf("❌ Failed to connect to the database: %v", err)
    }

    if err := DB.Ping(); err != nil {
        log.Fatalf("❌ Database ping failed: %v", err)
    }

    log.Println("✅ Connected to the SQLite database successfully!")
}`
		}
	} else if orm == "Ent" {
		if database == "Postgres" {
			return `package database

import (
    "context"
    "fmt"
    "log"
    "os"
    "entgo.io/ent/dialect"
    "entgo.io/ent/dialect/sql"
)

var DB *ent.Client

func Connect() {
    dbHost := getEnv("DB_HOST", "localhost")
    dbPort := getEnv("DB_PORT", "5432")
    dbUser := getEnv("DB_USER", "postgres")
    dbPassword := getEnv("DB_PASSWORD", "password")
    dbName := getEnv("DB_NAME", "mydb")

    dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
        dbHost, dbPort, dbUser, dbPassword, dbName)

    client, err := ent.Open("postgres", dsn)
    if err != nil {
        log.Fatalf("❌ Failed to connect to the database: %v", err)
    }
    DB = client

    if err := DB.Schema.Create(context.Background()); err != nil {
        log.Fatalf("❌ Failed to create schema: %v", err)
    }

    log.Println("✅ Connected to the PostgreSQL database successfully!")
}`
		}
		if database == "MySQL" {
			return `package database

import (
    "context"
    "fmt"
    "log"
    "os"
    "entgo.io/ent/dialect"
    "entgo.io/ent/dialect/sql"
)

var DB *ent.Client

func Connect() {
    dbHost := getEnv("DB_HOST", "localhost")
    dbPort := getEnv("DB_PORT", "3306")
    dbUser := getEnv("DB_USER", "root")
    dbPassword := getEnv("DB_PASSWORD", "password")
    dbName := getEnv("DB_NAME", "mydb")

    dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=True",
        dbUser, dbPassword, dbHost, dbPort, dbName)

    client, err := ent.Open("mysql", dsn)
    if err != nil {
        log.Fatalf("❌ Failed to connect to the database: %v", err)
    }
    DB = client

    if err := DB.Schema.Create(context.Background()); err != nil {
        log.Fatalf("❌ Failed to create schema: %v", err)
    }

    log.Println("✅ Connected to the MySQL database successfully!")
}`
		}
		if database == "SQLite" {
			return `package database

import (
    "context"
    "log"
    "os"
    "entgo.io/ent/dialect"
    "entgo.io/ent/dialect/sql"
)

var DB *ent.Client

func Connect() {
    dbName := getEnv("DB_NAME", "mydb.db")

    client, err := ent.Open("sqlite3", dbName)
    if err != nil {
        log.Fatalf("❌ Failed to connect to the database: %v", err)
    }
    DB = client

    if err := DB.Schema.Create(context.Background()); err != nil {
        log.Fatalf("❌ Failed to create schema: %v", err)
    }

    log.Println("✅ Connected to the SQLite database successfully!")
}`
		}
	}
	return "None"
}
