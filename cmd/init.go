package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/AlecAivazis/survey/v2"
	"github.com/spf13/cobra"
)

// initCmd represents the init command
var InitCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize a new Go web application",
	Run: func(cmd *cobra.Command, args []string) {
		var projectName, backend, database, orm string

		survey.AskOne(&survey.Input{Message: "What is your project name?"}, &projectName)
		survey.AskOne(&survey.Select{
			Message: "Choose your web framework:",
			Options: []string{"Fiber", "Gin"},
		}, &backend)
		survey.AskOne(&survey.Select{
			Message: "Choose your database system:",
			Options: []string{"Postgres", "MySQL", "SQLite"},
		}, &database)

		// Set default value for orm
		orm = "none"

		var useORM bool
		survey.AskOne(&survey.Confirm{
			Message: "Would you like to use an ORM?",
		}, &useORM)

		if useORM {
			survey.AskOne(&survey.Select{
				Message: "Choose your ORM framework:",
				Options: []string{"GORM", "XORM", "Ent"},
			}, &orm)
		}

		projectPath := filepath.Join(".", projectName)
		os.MkdirAll(projectPath, os.ModePerm)

		// Backend
		BackendCmd.Run(cmd, []string{projectPath, backend, database, orm})

		fmt.Println("✅ Project initialized successfully!")
	},
}
