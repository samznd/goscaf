package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/samznd/goweb/pkg/utils"

	"github.com/spf13/cobra"
)

// backendCmd represents the backend command
var BackendCmd = &cobra.Command{
	Use:   "backend",
	Short: "Generate server boilerplate",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 2 {
			fmt.Println("❌ Error: Missing arguments")
			return
		}

		projectPath, backend, database, orm := args[0], args[1], args[2], args[3]
		directories := []string{
			"cmd", "config", "internal", "internal/database", "internal/middleware",
			"internal/models", "internal/repositories", "internal/services",
			"internal/handlers", "internal/routes", "pkg", "scripts",
		}

		for _, dir := range directories {
			os.MkdirAll(projectPath+"/"+dir, 0755)
		}

		os.MkdirAll(projectPath, os.ModePerm)

		mainContent := getMainFile(backend, projectPath)
		databaseContent := getDatabaseFile(database, orm)
		envContent := `DB_USER=postgres
					   DB_PASSWORD=postgres
					   DB_HOST=localhost
					   DB_PORT=5432`

		utils.CreateFile(filepath.Join(projectPath+"/cmd/", "main.go"), mainContent)
		utils.CreateFile(filepath.Join(projectPath+"/config/", "database.go"), databaseContent)
		utils.CreateFile(filepath.Join(projectPath, ".env"), envContent)

		fmt.Println("✅ Backend initialized!")
		installDependencies(projectPath, backend)

	},
}

func runCommand(dir, command string) {
	cmd := exec.Command("sh", "-c", command)
	cmd.Dir = dir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Printf("❌ Error executing command: %s\n", command)
	}
}

func installDependencies(projectName, backend string) {
	fmt.Println("📦 Initializing Go module...")
	runCommand(projectName, "go mod init "+projectName)

	fmt.Println("📦 Installing dependencies...")
	if backend == "fiber" {
		runCommand(projectName, "go get github.com/gofiber/fiber/v2")
	} else {
		runCommand(projectName, "go get github.com/gin-gonic/gin")
	}
	runCommand(projectName, "go get github.com/lib/pq")

	fmt.Println("✅ Dependencies installed successfully!")
}
