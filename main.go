package main

import (
	"fmt"
	"os"

	"github.com/samznd/goscaf/internal/generator"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "goscaf",
	Short: "A CLI to generate Go web application boilerplate",
}

func main() {
	RootCmd.SetHelpCommand(&cobra.Command{Hidden: true})
	RootCmd.SetFlagErrorFunc(func(cmd *cobra.Command, err error) error {
		fmt.Printf("❌ %v\n\n", err)
		cmd.Usage()
		os.Exit(1)
		return nil
	})

	RootCmd.AddCommand(generator.InitCmd)

	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
