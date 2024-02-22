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

// Copy file function
func copyFile(src, dst string) error {
	srcFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	srcFileInfo, err := srcFile.Stat()
	if err != nil {
		return err
	}

	if !srcFileInfo.Mode().IsRegular() {
		return fmt.Errorf("%s is not a regular file", src)
	}

	dstFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer dstFile.Close()

	_, err = io.Copy(dstFile, srcFile)
	if err != nil {
		return err
	}

	return nil
}

// Remove contents in directory function
func removeContentsInDirectory(directoryPath string) error {
	entries, err := os.ReadDir(directoryPath)
	if err != nil {
		return err
	}

	for _, entry := range entries {
		entryPath := filepath.Join(directoryPath, entry.Name())

		if err := os.RemoveAll(entryPath); err != nil {
			return err
		}
	}

	return nil
}

func main() {
	pathToJsons := "bo4e_schemas/bo"
	tempDir := "temp/bo"
	finalDir := "destination/bo"

	// Open the directory
	dir, err := os.Open(pathToJsons)
	if err != nil {
		fmt.Println("Error opening directory:", err)
		return
	}
	defer dir.Close()

	jsonEntries, err := dir.ReadDir(-1)
	if err != nil {
		fmt.Println("error reading directory:")
		return
	}

	//Iterate over the json schemas
	for _, jsonEntry := range jsonEntries {

		jsonFileName := filepath.Base(jsonEntry.Name())
		// name of the file without extension
		fileName := strings.TrimSuffix(jsonFileName, ".json")
		// go file name
		goFileName := fileName + ".go"
		// make paths to jsons and go files in tempDir and finalDir
		pathToJson := strings.Replace(filepath.Join(pathToJsons, jsonFileName), "\\", "/", -1)
		tempFilePath := strings.Replace(filepath.Join(tempDir, goFileName), "\\", "/", -1)
		finalFilePath := strings.Replace(filepath.Join(finalDir, goFileName), "\\", "/", -1)

		//Run quicktype
		err := runQuicktype(pathToJson, tempFilePath)
		if err != nil {
			fmt.Println(err)
			return
		}
		// Copy the generated go file to the final directory
		if err := copyFile(tempFilePath, finalFilePath); err != nil {
			return
		}
	}

	// Remove contents in the temp/bo directory
	err = removeContentsInDirectory(tempDir)
	if err != nil {
		fmt.Println("Error removing contents in temp/bo directory:", err)
		return
	}

}
