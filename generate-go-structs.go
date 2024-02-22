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
	//finalDir := "destination/bo"

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
		fmt.Println("jsonFileName ", jsonFileName)
		fileName := strings.TrimSuffix(jsonFileName, ".json")
		fmt.Println("fileName ", fileName)
		goFileName := fileName + ".go"
		// check if the file with this name exists in a temp directory
		pathToJson := strings.Replace(filepath.Join(pathToJsons, jsonFileName), "\\", "/", -1)
		tempFilePath := strings.Replace(filepath.Join(tempDir, goFileName), "\\", "/", -1)
		fmt.Println("pathToJson ", pathToJson)
		fmt.Println("tempFilePath ", tempFilePath)
		//if _, err := os.Stat(tempFilePath); os.IsNotExist(err) {
		//	// File does not exist, create it
		//	os.Create(tempFilePath)
		//	fmt.Println("File", tempFilePath, "created successfully.")
		//} else if err != nil {
		//	// Some other error occurred
		//	fmt.Println("Error:", err)
		//	return
		//}
		//Run quicktype
		err := runQuicktype(pathToJson, tempFilePath)
		if err != nil {
			fmt.Println(err)
			return
		}
	}



			//// Copy the generated Go file to the final destination
			//err = copyGoFileToDestination(tempDir, finalDir, fileName)
			//if err != nil {
			//	fmt.Println("Error copying Go file to destination:", err)
			//	return nil
			//}

			//// Remove contents in the temp/bo directory
			//err = removeContentsInDirectory(tempDir)
			//if err != nil {
			//	fmt.Println("Error removing contents in temp/bo directory:", err)
			//	return nil
			//}

			//fmt.Printf("Go files generated and copied successfully for %s\n", fileName)

		}
		return
	}



}
