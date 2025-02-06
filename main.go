package main

import (
	"fmt"
	"os"

	"github.com/samznd/goscaf/cmd"

	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "goscaf",
	Short: "A CLI to generate Go web application boilerplate",
}

func main() {
	RootCmd.AddCommand(cmd.InitCmd)
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
