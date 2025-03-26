package cmd

import (
	"fmt"
	"strings"
)

func getMainFile(backend string, projectName string) string {
	switch strings.ToLower(backend) {
	case "fiber":
		return fmt.Sprintf(`package main

import (
    "log"
    "github.com/gofiber/fiber/v2"
    "%s/config"
    "%s/pkg/utils"
)

func main() {
    utils.InitialEnv()
    config.Connect()
    app := fiber.New()

    // Define routes
    app.Get("/ping", func(c *fiber.Ctx) error {
        return c.JSON(fiber.Map{"message": "pong"})
    })

    log.Println("🚀 Fiber server is running on http://localhost:3000")
    app.Listen(":3000")
}`, projectName, projectName)
	case "gin":
		return fmt.Sprintf(`package main

import (
    "log"
    "net/http"
    "github.com/gin-gonic/gin"
    "%s/config"
    "%s/pkg/utils"
)

func main() {
    utils.InitialEnv()
    config.Connect()

    r := gin.Default()

    // Define routes
    r.GET("/ping", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{"message": "pong"})
    })

    log.Println("🚀 Gin server is running on http://localhost:3000")
    r.Run(":3000")
}`, projectName, projectName)
	case "echo":
		return fmt.Sprintf(`package main

import (
    "net/http"
    "github.com/labstack/echo/v4"
    "%s/config"
    "%s/pkg/utils"
)

func main() {
    utils.InitialEnv()
    config.Connect()

    e := echo.New()

    // Define routes
    e.GET("/ping", func(c echo.Context) error {
        return c.JSON(http.StatusOK, map[string]string{"message": "pong"})
    })

    e.Logger.Fatal(e.Start(":3000"))
}`, projectName, projectName)
	case "chi":
		return fmt.Sprintf(`package main

import (
    "net/http"
    "github.com/go-chi/chi"
    "github.com/go-chi/chi/middleware"
    "%s/config"
    "%s/pkg/utils"
)

func main() {
    utils.InitialEnv()
    config.Connect()

    r := chi.NewRouter()
    r.Use(middleware.Logger)

    // Define routes
    r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusOK)
        w.Write([]byte(`+"`{\"message\": \"pong\"}`"+`))
    })

    http.ListenAndServe(":3000", r)
}`, projectName, projectName)
	case "iris":
		return fmt.Sprintf(`package main

import (
    "github.com/kataras/iris/v12"
    "%s/config"
    "%s/pkg/utils"
)

func main() {
    utils.InitialEnv()
    config.Connect()

    app := iris.New()

    // Define routes
    app.Get("/ping", func(ctx iris.Context) {
        ctx.JSON(iris.Map{"message": "pong"})
    })

    app.Listen(":3000")
}`, projectName, projectName)
	default:
		return ""
	}
}

func getDatabaseFile(database string, orm string, projectName string) string {
	if strings.ToLower(orm) != "none" {
		return SetupORM(orm, database)
	}

	template := ""
	switch strings.ToLower(database) {
	case "postgres":
		template = `package config

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
    dbPort := os.Getenv("DB_PORT", "5432")
    dbUser := os.Getenv("DB_USER", "postgres")
    dbPassword := os.Getenv("DB_PASSWORD")
    dbName := os.Getenv("DB_NAME", "mydb")

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
}`
	case "mysql":
		template = `package config

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
    dbName := os.Getenv("DB_NAME", "mydb")

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
	case "sqlite":
		template = `package config

import (
    "database/sql"
    "log"
    "os"
    _ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func Connect() {
    dbName := os.Getenv("DB_NAME", "mydb")

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
}`
	default:
		fmt.Printf("Warning: Unknown database type '%s'\n", database)
		return "None"
	}

	return template
}

func SetupORM(orm string, database string) string {
	ormLower := strings.ToLower(orm)
	dbLower := strings.ToLower(database)

	if ormLower == "gorm" {
		switch dbLower {
		case "postgres":
			return `package config

import (
    "fmt"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
    "log"
)

var DB *gorm.DB

func Connect() {
    dbHost := os.Getenv("DB_HOST")
    dbPort := os.Getenv("DB_PORT", "5432")
    dbUser := os.Getenv("DB_USER", "postgres")
    dbPassword := os.Getenv("DB_PASSWORD")
    dbName := os.Getenv("DB_NAME", "mydb")

    dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
        dbHost, dbPort, dbUser, dbPassword, dbName)

    var err error
    DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatalf("❌ Failed to connect to the database: %v", err)
    }

    log.Println("✅ Connected to the PostgreSQL database successfully!")
}`
		case "mysql":
			return `package config

import (
    "fmt"
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
    "log"
)

var DB *gorm.DB

func Connect() {
    dbHost := os.Getenv("DB_HOST")
    dbPort := os.Getenv("DB_PORT")
    dbUser := os.Getenv("DB_USER")
    dbPassword := os.Getenv("DB_PASSWORD")
    dbName := os.Getenv("DB_NAME", "mydb")

    dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
        dbUser, dbPassword, dbHost, dbPort, dbName)

    var err error
    DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatalf("❌ Failed to connect to the database: %v", err)
    }

    log.Println("✅ Connected to the MySQL database successfully!")
}`
		case "sqlite":
			return `package config

import (
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
    "log"
    "os"
)

var DB *gorm.DB

func Connect() {
    dbName := os.Getenv("DB_NAME", "mydb.db")

    var err error
    DB, err = gorm.Open(sqlite.Open(dbName), &gorm.Config{})
    if err != nil {
        log.Fatalf("❌ Failed to connect to the database: %v", err)
    }

    log.Println("✅ Connected to the SQLite database successfully!")
}`
		}
	} else if ormLower == "xorm" {
		switch dbLower {
		case "postgres":
			return `package config

import (
    "fmt"
    "log"
    "os"
    "xorm.io/xorm"
    _ "github.com/lib/pq"
)

var DB *xorm.Engine

func Connect() {
    dbHost := os.Getenv("DB_HOST")
    dbPort := os.Getenv("DB_PORT", "5432")
    dbUser := os.Getenv("DB_USER", "postgres")
    dbPassword := os.Getenv("DB_PASSWORD")
    dbName := os.Getenv("DB_NAME", "mydb")

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
		case "mysql":
			return `package config

import (
    "fmt"
    "log"
    "os"
    "xorm.io/xorm"
    _ "github.com/go-sql-driver/mysql"
)

var DB *xorm.Engine

func Connect() {
    dbHost := os.Getenv("DB_HOST")
    dbPort := os.Getenv("DB_PORT")
    dbUser := os.Getenv("DB_USER")
    dbPassword := os.Getenv("DB_PASSWORD")
    dbName := os.Getenv("DB_NAME", "mydb")

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
		case "sqlite":
			return `package config

import (
    "log"
    "os"
    "xorm.io/xorm"
    _ "github.com/mattn/go-sqlite3"
)

var DB *xorm.Engine

func Connect() {
    dbName := os.Getenv("DB_NAME")

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
	} else if ormLower == "ent" {
		switch dbLower {
		case "postgres":
			return `package config

import (
    "context"
    "fmt"
    "log"
    "entgo.io/ent/dialect"
    "entgo.io/ent/dialect/sql"
)

var DB *ent.Client

func Connect() {
    dbHost := os.Getenv("DB_HOST")
    dbPort := os.Getenv("DB_PORT", "5432")
    dbUser := os.Getenv("DB_USER", "postgres")
    dbPassword := os.Getenv("DB_PASSWORD")
    dbName := os.Getenv("DB_NAME", "mydb")

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
		case "mysql":
			return `package config

import (
    "context"
    "fmt"
    "log"
    "entgo.io/ent/dialect"
    "entgo.io/ent/dialect/sql"
)

var DB *ent.Client

func Connect() {
    dbHost := os.Getenv("DB_HOST")
    dbPort := os.Getenv("DB_PORT")
    dbUser := os.Getenv("DB_USER")
    dbPassword := os.Getenv("DB_PASSWORD")
    dbName := os.Getenv("DB_NAME", "mydb")

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
		case "sqlite":
			return `package config

import (
    "context"
    "log"
    "entgo.io/ent/dialect"
    "entgo.io/ent/dialect/sql"
)

var DB *ent.Client

func Connect() {
    dbName := os.Getenv("DB_NAME", "mydb.db")

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

func getUtilsFile() string {
	return `package utils

import (
	"log"

	"github.com/joho/godotenv"
)

func InitialEnv() error {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	return nil
}`
}

func getDockerFile(projectName string) string {
	return fmt.Sprintf(`FROM golang:1.17-alpine AS builder

WORKDIR /app

# Copy the Go Modules manifests
COPY go.mod go.sum ./

# Download the Go module dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the Go application
RUN go build -o %s .

# Use a minimal base image
FROM alpine:latest

WORKDIR /root/

# Copy the built application from the builder stage
COPY --from=builder /app/%s .

# Expose the application port
EXPOSE 8080

# Run the application
CMD ["./%s"]

# Note: Customize the Dockerfile as needed for your specific project requirements.
# For example, you may need to add environment variables, additional dependencies, or other configurations.
`, projectName, projectName, projectName)
}

func getDockerComposeFile(database string) string {
	var dbService string

	switch database {
	case "Postgres":
		dbService = `
  db:
    image: postgres:latest
    environment:
      POSTGRES_USER: user
      POSTGRES_PASSWORD: password
      POSTGRES_DB: mydb
    ports:
      - "5432:5432"
    volumes:
      - db-data:/var/lib/postgresql/data
`
	case "MySQL":
		dbService = `
  db:
    image: mysql:latest
    environment:
      MYSQL_ROOT_PASSWORD: password
      MYSQL_DATABASE: mydb
      MYSQL_USER: user
      MYSQL_PASSWORD: password
    ports:
      - "3306:3306"
    volumes:
      - db-data:/var/lib/mysql
`
	case "SQLite":
		dbService = `
  db:
    image: nouchka/sqlite3
    volumes:
      - db-data:/data
`

	default:
		dbService = ""
	}

	return fmt.Sprintf(`version: '3.8'

services:
  app:
    build: .
    ports:
      - "8080:8080"
    depends_on:
      - db
    environment:
      - DB_HOST=db
      - DB_PORT=5432
      - DB_USER=user
      - DB_PASSWORD=password
      - DB_NAME=mydb
%s

volumes:
  db-data:
`, dbService)
}
