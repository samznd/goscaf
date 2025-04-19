package utils

import (
	"fmt"
	"os"
	"path/filepath"
)

// CreateFile writes content into a file at the given path
func CreateFile(filePath, content string) error {
	// Create the file
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	// Write content
	_, err = file.WriteString(content)
	if err != nil {
		return err
	}

	fmt.Println("✅ Created file:", filePath)
	return nil
}

func CreateTemplate(dir, filename, content, projectPath string) {
	fullPath := filepath.Join(projectPath, "internal", dir, filename)
	if err := CreateFile(fullPath, content); err != nil {
		fmt.Printf("❌ Error creating %s: %v\n", filename, err)
		os.Exit(1)
		os.RemoveAll(fullPath)
	}
}
