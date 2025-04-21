package templates

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/samznd/goscaf/pkg/utils"
	"github.com/spf13/cobra"
)

var InitTemplateCmd = &cobra.Command{
	Use:   "backend_templates",
	Short: "Generate backend template files based on frameworks, orm, etc",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 3 {
			fmt.Println("❌ Error: Missing arguments")
			fmt.Printf("Received args: %v\n", args)
			os.Exit(1)
		}
		projectName, backend, orm := args[0], args[1], args[2]

		repositoryTemplate := RepositoryTemplate(strings.ToLower(orm))
		serviceTemplate := ServiceTemplate(projectName)
		handlerTemplate := HandlerGenerator(projectName, strings.ToLower(backend), strings.ToLower(orm))
		setupRoutesTemplate := SetupRoutesTemplate(projectName, strings.ToLower(backend))
		projectPath := filepath.Join(".", projectName)
		utils.CreateTemplate("repositories", "repository.go", repositoryTemplate, projectPath)
		utils.CreateTemplate("services", "service.go", serviceTemplate, projectPath)
		utils.CreateTemplate("handlers", "handler.go", handlerTemplate, projectPath)
		utils.CreateTemplate("routes", "routes.go", setupRoutesTemplate, projectPath)

	},
}
