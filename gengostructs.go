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

//// Remove contents in directory function
//func removeContentsInDirectory(directoryPath string) error {
//	entries, err := os.ReadDir(directoryPath)
//	if err != nil {
//		return err
//	}
//
//	for _, entry := range entries {
//		entryPath := filepath.Join(directoryPath, entry.Name())
//
//		if err := os.RemoveAll(entryPath); err != nil {
//			return err
//		}
//	}
//
//	return nil
//}

func main() {
	root := "bo4e_schemas"
	tempDir := "temp"
	finalDir := "final"

	// Create a temporary directory with 0777 permission (full access)
	errTemp := os.Mkdir(tempDir, 0777)
	if errTemp != nil {
		fmt.Println("Error creating temporary directory:", errTemp)
		return
	}

	// Create a final directory with 0777 permission (full access)
	errFinal := os.Mkdir(finalDir, 0777)
	if errFinal != nil {
		fmt.Println("Error creating final temporary directory:", errFinal)
		return
	}

	walkFunc := func(path string, info os.FileInfo, err error) error {
		// check if there was an error accessing the file or dir
		if err != nil {
			fmt.Printf("Error accessing %s %v", path, err)
			return nil // skip this file or dir
		}

		// check if it is a regular file
		if !info.IsDir() {
			// get the file name
			// Calculate the corresponding path in the temporary directory
			pathToFile := filepath.Join(tempDir, path[len(root):])
			pathWithoutExt := strings.TrimSuffix(pathToFile, ".json")
			goPathToFile := pathWithoutExt + ".go"
			// Run quicktype
			err := runQuicktype(path, goPathToFile)
			if err != nil {
				fmt.Println(err)
				return nil
			}
			fmt.Println("goPathToFile:", goPathToFile)

		}
		return nil

	}

	// Traverse the directory tree using Walk
	err := filepath.Walk(root, walkFunc)
	if err != nil {
		fmt.Printf("Error walking the path %s: %v\n", root, err)
	}

	//// Extract the file name without extension
	//jsonFileName := strings.TrimSuffix(filepath.Base(jsonPath), filepath.Ext(jsonPath))
	//
	//// Run quicktype
	//err := runQuicktype(jsonPath, temporaryDestination)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//
	//// Trim "/Angebot.go" from temporaryDestination
	//temporaryDest := strings.TrimSuffix(temporaryDestination, "/Angebot.go")
	//
	//// Copy the generated Go file to the final destination
	//err = copyGoFileToDestination(temporaryDest, finalDestination, jsonFileName)
	//if err != nil {
	//	fmt.Println("Error copying Go file to destination:", err)
	//	return
	//}
	//
	////// Remove contents in the temp/bo directory
	////err = removeContentsInDirectory(temporaryDest)
	////if err != nil {
	////	fmt.Println("Error removing contents in temp/bo directory:", err)
	////	return
	////}
	//
	//fmt.Printf("Go files generated and copied successfully for %s\n", jsonFileName)
}
