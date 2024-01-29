package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func copyDirStructure(source, destination string) error {
	err := filepath.Walk(source, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		relPath := strings.Replace(path, source, "", 1)
		destPath := filepath.Join(destination, relPath)

		if info.IsDir() {
			if err := os.MkdirAll(destPath, 0755); err != nil {
				return err
			}
		}
		return nil
	})

	return err
}

func main() {
	source := "./bo4e_schemas"
	destination := "./destination"

	err := copyDirStructure(source, destination)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Directory structure copied successfully.")
	}
}
