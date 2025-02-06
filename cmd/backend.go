package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/samznd/goscaf/pkg/utils"

	"github.com/spf13/cobra"
)

// backendCmd represents the backend command
var BackendCmd = &cobra.Command{
	Use:   "backend",
	Short: "Generate backend directories and files for a Go web application",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 4 {
			fmt.Println("❌ Error: Missing arguments")
			fmt.Printf("Received args: %v\n", args)
			os.Exit(1)
		}

		projectPath, backend, database, orm := args[0], args[1], args[2], args[3]

		// Create directories
		directories := []string{
			"cmd", "config", "internal", "internal/middleware",
			"internal/models", "internal/repositories", "internal/services",
			"internal/handlers", "internal/routes", "pkg/utils", "scripts",
		}

		for _, dir := range directories {
			fullPath := filepath.Join(projectPath, dir)
			if err := os.MkdirAll(fullPath, 0755); err != nil {
				fmt.Printf("Error creating directory %s: %v\n", fullPath, err)
			}
		}

		// Generate files
		mainContent := getMainFile(backend, projectPath)
		databaseContent := getDatabaseFile(database, orm)

		if databaseContent == "None" {
			fmt.Printf("Error: Invalid database configuration. Database: %s, ORM: %s\n", database, orm)
			os.Exit(1)
		}

		envContent := `DB_USER=postgres
DB_PASSWORD=postgres
DB_HOST=localhost
DB_PORT=5432
DB_NAME=mydb`

		utilsContent := getUtilsFile()
		dockerfileContent := getDockerFile(projectPath)
		dockerComposeContent := getDockerComposeFile(database)
		// Create files
		if err := utils.CreateFile(filepath.Join(projectPath, "cmd", "main.go"), mainContent); err != nil {
			fmt.Printf("Error creating main.go: %v\n", err)
		}
		if err := utils.CreateFile(filepath.Join(projectPath, "config", "database.go"), databaseContent); err != nil {
			fmt.Printf("Error creating database.go: %v\n", err)
		}
		if err := utils.CreateFile(filepath.Join(projectPath, ".env"), envContent); err != nil {
			fmt.Printf("Error creating .env: %v\n", err)
		}
		if err := utils.CreateFile(filepath.Join(projectPath, "pkg/utils", "env_utils.go"), utilsContent); err != nil {
			fmt.Printf("Error creating env_utils.go: %v\n", err)
		}
		if err := utils.CreateFile(filepath.Join(projectPath, "Dockerfile"), dockerfileContent); err != nil {
			fmt.Printf("Error creating Dockerfile: %v\n", err)
		}
		if err := utils.CreateFile(filepath.Join(projectPath, "docker-compose.yml"), dockerComposeContent); err != nil {
			fmt.Printf("Error creating docker-compose.yml: %v\n", err)
		}

		// Create database migration script
		// Initialize go.mod and install dependencies
		installDependencies(projectPath, backend, database, orm)
	},
}

func installDependencies(projectPath, backend string, database string, orm string) {
	fmt.Println("📦 Initializing Go module...")
	runCommand(projectPath, "go mod init "+projectPath)

	fmt.Println("📦 Installing dependencies...")

	// Install backend framework
	switch strings.ToLower(backend) {
	case "fiber":
		runCommand(projectPath, "go get github.com/gofiber/fiber/v2")
	case "gin":
		runCommand(projectPath, "go get github.com/gin-gonic/gin")
	case "echo":
		runCommand(projectPath, "go get github.com/labstack/echo/v4")
	case "chi":
		runCommand(projectPath, "go get github.com/go-chi/chi")
	case "iris":
		runCommand(projectPath, "go get github.com/kataras/iris/v12")

	default:
		fmt.Printf("❌ Error: Invalid backend: %s\n", backend)
		os.Exit(1)
	}

	// Install database driver
	switch strings.ToLower(database) {
	case "postgres":
		runCommand(projectPath, "go get github.com/lib/pq")
	case "mysql":
		runCommand(projectPath, "go get github.com/go-sql-driver/mysql")
	case "sqlite":
		runCommand(projectPath, "go get github.com/mattn/go-sqlite3")

	default:
		fmt.Printf("❌ Error: Invalid database: %s\n", database)
		os.Exit(1)
	}

	// Install ORM if selected
	switch strings.ToLower(orm) {
	case "gorm":
		runCommand(projectPath, "go get gorm.io/gorm")
		// Install GORM database drivers
		switch strings.ToLower(database) {
		case "postgres":
			runCommand(projectPath, "go get gorm.io/driver/postgres")
		case "mysql":
			runCommand(projectPath, "go get gorm.io/driver/mysql")
		case "sqlite":
			runCommand(projectPath, "go get gorm.io/driver/sqlite")

		default:
			fmt.Printf("❌ Error: Invalid database: %s\n", database)
			os.Exit(1)
		}
	case "xorm":
		runCommand(projectPath, "go get xorm.io/xorm")
	case "ent":
		runCommand(projectPath, "go get entgo.io/ent")
		runCommand(projectPath, "go get entgo.io/ent/cmd/ent")
	default:
		fmt.Printf("❌ Error: Invalid ORM: %s\n", orm)
		os.Exit(1)
	}

	// Install common utilities
	runCommand(projectPath, "go get github.com/joho/godotenv")
	runCommand(projectPath, "go get golang.org/x/crypto")

	// Fix missing dependencies
	runCommand(projectPath, "go get github.com/mattn/go-isatty@v0.0.20")

	// Tidy up modules and ensure all dependencies are properly downloaded
	runCommand(projectPath, "go mod tidy")
	runCommand(projectPath, "go mod download")

	fmt.Println("✅ Dependencies installed successfully!")
}

func runCommand(dir, command string) {
	cmd := exec.Command("sh", "-c", command)
	cmd.Dir = dir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Printf("❌ Error executing command '%s': %v\n", command, err)
		os.Exit(1)
	}
}
