package main

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func runQuicktype(jsonPath, outputDir string) error {
	cmd := exec.Command("quicktype", "--src", jsonPath, "-o", outputDir, "--src-lang", "schema", "--lang", "go", "--multi-file-output")

	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("error running quicktype: %s\nOutput:\n%s", err, output)
	}
	return nil
}

func copyGoFileToDestination(temporaryPath, destinationPath, jsonFileName string) error {
	// Assuming that the generated Go file has the same name as the JSON file
	goFileName := jsonFileName + ".go"

	// Create the destination directory if it doesn't exist
	if err := os.MkdirAll(destinationPath, os.ModePerm); err != nil {
		return err
	}

	// Calculate paths
	temporaryFilePath := filepath.Join(temporaryPath, goFileName)
	destinationFilePath := filepath.Join(destinationPath, goFileName)

	// Copy the Go file to the destination
	if err := copyFile(temporaryFilePath, destinationFilePath); err != nil {
		return err
	}

	return nil
}

// Copy file function
func copyFile(src, dest string) error {
	sourceFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	destFile, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer destFile.Close()

	_, err = io.Copy(destFile, sourceFile)
	if err != nil {
		return err
	}

	return destFile.Sync()
}

func main() {
	jsonPath := "bo4e_schemas/bo/Angebot.json"
	temporaryDestination := "temp/bo/Angebot.go"
	finalDestination := "destination"

	// Extract the file name without extension
	jsonFileName := strings.TrimSuffix(filepath.Base(jsonPath), filepath.Ext(jsonPath))

	// Run quicktype
	err := runQuicktype(jsonPath, temporaryDestination)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Copy the generated Go file to the final destination
	err = copyGoFileToDestination(temporaryDestination, finalDestination, jsonFileName)
	if err != nil {
		fmt.Println("Error copying Go file to destination:", err)
		return
	}

	fmt.Printf("Go files generated and copied successfully for %s\n", jsonFileName)
}
