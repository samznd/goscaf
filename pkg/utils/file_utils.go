package utils

import (
	"fmt"
	"os"
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
