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
)

func main() {
    config.Connect()
    app := fiber.New()

    // Define routes
    app.Get("/ping", func(c *fiber.Ctx) error {
        return c.JSON(fiber.Map{"message": "pong"})
    })

    log.Println("🚀 Fiber server is running on http://localhost:3000")
    app.Listen(":3000")
}`, projectName)
	case "gin":
		return fmt.Sprintf(`package main

import (
    "fmt"
    "log"
    "net/http"
    "github.com/gin-gonic/gin"
    "%s/config"
)

func main() {
    config.Connect()

    r := gin.Default()

    // Define routes
    r.GET("/ping", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{"message": "pong"})
    })

    log.Println("🚀 Gin server is running on http://localhost:3000")
    r.Run(":3000")
}`, projectName)
	case "echo":
		return fmt.Sprintf(`package main

import (
    "net/http"
    "github.com/labstack/echo/v4"
    "%s/config"
)

func main() {
    config.Connect()

    e := echo.New()

    // Define routes
    e.GET("/ping", func(c echo.Context) error {
        return c.JSON(http.StatusOK, map[string]string{"message": "pong"})
    })

    e.Logger.Fatal(e.Start(":3000"))
}`, projectName)
	case "chi":
		return fmt.Sprintf(`package main

import (
    "net/http"
    "github.com/go-chi/chi"
    "github.com/go-chi/chi/middleware"
    "%s/config"
)

func main() {
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
}`, projectName)
	case "iris":
		return fmt.Sprintf(`package main

import (
    "github.com/kataras/iris/v12"
    "%s/config"
)

func main() {
    config.Connect()

    app := iris.New()

    // Define routes
    app.Get("/ping", func(ctx iris.Context) {
        ctx.JSON(iris.Map{"message": "pong"})
    })

    app.Listen(":3000")
}`, projectName)
	default:
		return ""
	}
}

func getDatabaseFile(database string, orm string) string {
	if strings.ToLower(orm) != "none" {
		return SetupORM(orm, database)
	}

	switch strings.ToLower(database) {
	case "postgres":
		return `package config

import (
    "database/sql"
    "fmt"
    "log"
    "test/pkg/utils"
    _ "github.com/lib/pq"
)

var DB *sql.DB

func Connect() {
    dbHost := utils.GetEnv("DB_HOST", "localhost")
    dbPort := utils.GetEnv("DB_PORT", "5432")
    dbUser := utils.GetEnv("DB_USER", "postgres")
    dbPassword := utils.GetEnv("DB_PASSWORD", "password")
    dbName := utils.GetEnv("DB_NAME", "mydb")

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
`
	case "mysql":
		return `package config

import (
    "database/sql"
    "fmt"
    "log"   
    "github.com/go-sql-driver/mysql"
)
    
var DB *sql.DB

func Connect() {
    dbHost := utils.GetEnv("DB_HOST", "localhost")
    dbPort := utils.GetEnv("DB_PORT", "3306")
    dbUser := utils.GetEnv("DB_USER", "root")
    dbPassword := utils.GetEnv("DB_PASSWORD", "password")
    dbName := utils.GetEnv("DB_NAME", "mydb")

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
		return `package config

import (
    "database/sql"
    "log"
    _ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB


func Connect() {
    dbName := utils.GetEnv("DB_NAME", "mydb")

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
    dbHost := utils.GetEnv("DB_HOST", "localhost")
    dbPort := utils.GetEnv("DB_PORT", "5432")
    dbUser := utils.GetEnv("DB_USER", "postgres")
    dbPassword := utils.GetEnv("DB_PASSWORD", "password")
    dbName := utils.GetEnv("DB_NAME", "mydb")

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
    dbHost := utils.GetEnv("DB_HOST", "localhost")
    dbPort := utils.GetEnv("DB_PORT", "3306")
    dbUser := utils.GetEnv("DB_USER", "root")
    dbPassword := utils.GetEnv("DB_PASSWORD", "password")
    dbName := utils.GetEnv("DB_NAME", "mydb")

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
)

var DB *gorm.DB

func Connect() {
    dbName := utils.GetEnv("DB_NAME", "mydb.db")

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
    "xorm.io/xorm"
    _ "github.com/lib/pq"
)

var DB *xorm.Engine

func Connect() {
    dbHost := utils.GetEnv("DB_HOST", "localhost")
    dbPort := utils.GetEnv("DB_PORT", "5432")
    dbUser := utils.GetEnv("DB_USER", "postgres")
    dbPassword := utils.GetEnv("DB_PASSWORD", "password")
    dbName := utils.GetEnv("DB_NAME", "mydb")

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
    "xorm.io/xorm"
    _ "github.com/go-sql-driver/mysql"
)

var DB *xorm.Engine

func Connect() {
    dbHost := utils.GetEnv("DB_HOST", "localhost")
    dbPort := utils.GetEnv("DB_PORT", "3306")
    dbUser := utils.GetEnv("DB_USER", "root")
    dbPassword := utils.GetEnv("DB_PASSWORD", "password")
    dbName := utils.GetEnv("DB_NAME", "mydb")

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
    "xorm.io/xorm"
    _ "github.com/mattn/go-sqlite3"
)

var DB *xorm.Engine

func Connect() {
    dbName := utils.GetEnv("DB_NAME", "mydb.db")

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
    dbHost := utils.GetEnv("DB_HOST", "localhost")
    dbPort := utils.GetEnv("DB_PORT", "5432")
    dbUser := utils.GetEnv("DB_USER", "postgres")
    dbPassword := utils.GetEnv("DB_PASSWORD", "password")
    dbName := utils.GetEnv("DB_NAME", "mydb")

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
    dbHost := utils.GetEnv("DB_HOST", "localhost")
    dbPort := utils.GetEnv("DB_PORT", "3306")
    dbUser := utils.GetEnv("DB_USER", "root")
    dbPassword := utils.GetEnv("DB_PASSWORD", "password")
    dbName := utils.GetEnv("DB_NAME", "mydb")

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
    dbName := utils.GetEnv("DB_NAME", "mydb.db")

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
    "os"
)

func GetEnv(key, fallback string) string {
    value, exists := os.LookupEnv(key)
    if !exists {
        return fallback
    }
    return value
}`
}
