package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func runQuicktype(jsonPath, outputFile string) error {
	cmd := exec.Command("quicktype", "--lang", "go", "--src", jsonPath, "--out", outputFile)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("error running quicktype: %s\nOutput:\n%s", err, output)
	}
	return nil
}

func main() {
	jsonPath := "./bo4e_schemas/bo/Angebot.json"
	temporaryDestination := "./temp"

	// Extract the file name without extension
	jsonFileName := strings.TrimSuffix(filepath.Base(jsonPath), filepath.Ext(jsonPath))

	// Create temporary directory if it doesn't exist
	if err := os.MkdirAll(temporaryDestination, os.ModePerm); err != nil {
		fmt.Println("Error creating temporary directory:", err)
		return
	}

	// Specify the output file path
	outputFilePath := filepath.Join(temporaryDestination, fmt.Sprintf("%s.go", jsonFileName))

	// Run quicktype
	err := runQuicktype(jsonPath, outputFilePath)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Go file generated successfully for %s\n", jsonFileName)
}
