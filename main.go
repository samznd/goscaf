package main

import (
	"fmt"
	"go-boilerplate-cli/cmd"
	"os"

	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "goweb",
	Short: "A CLI to generate Go web application boilerplate",
}

func main() {
	RootCmd.AddCommand(cmd.InitCmd)
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
