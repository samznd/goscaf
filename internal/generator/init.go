package generator

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

		err := survey.AskOne(&survey.Input{Message: "What is your project name?"}, &projectName)
		if err != nil {
			fmt.Println("\nOperation canceled by user.")
			os.Exit(1)
		}
		err = survey.AskOne(&survey.Select{
			Message: "Choose your web framework:",
			Options: []string{"Fiber", "Gin", "Echo", "Chi", "Iris"},
		}, &backend)
		if err != nil {
			fmt.Println("\nOperation canceled by user.")
			os.Exit(1)
		}
		err = survey.AskOne(&survey.Select{
			Message: "Choose your database system:",
			Options: []string{"Postgres", "MySQL", "SQLite"},
		}, &database)
		if err != nil {
			fmt.Println("\nOperation canceled by user.")
			os.Exit(1)
		}

		orm = "none"

		var useORM bool
		err = survey.AskOne(&survey.Confirm{Message: "Would you like to use an ORM?"}, &useORM)
		if err != nil {
			fmt.Println("\nOperation canceled by user.")
			os.Exit(1)
		}

		if useORM {
			err = survey.AskOne(&survey.Select{
				Message: "Choose your ORM framework:",
				Options: []string{"GORM", "XORM", "Ent", "SQLBoiler"},
			}, &orm)
			if err != nil {
				fmt.Println("\nOperation canceled by user.")
				os.Exit(1)
			}
		}

		projectPath := filepath.Join(".", projectName)
		os.MkdirAll(projectPath, os.ModePerm)

		ScaffoldBackendCmd.Run(cmd, []string{projectPath, backend, database, orm})

		fmt.Println("✅ Project initialized successfully!")
	},
}
