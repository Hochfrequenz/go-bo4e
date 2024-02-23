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
func createDir(dirPath string) error {
	if err := os.MkdirAll(dirPath, 0755); err != nil && !os.IsExist(err) {
		return err
	}
	fmt.Println("directory created successfully")
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
	sourcePath := "bo4e_schemas"
	tempDir := "temp"
	finalDir := "bo4e_structs"

	// Open the directory
	dir, err := os.Open(sourcePath)
	if err != nil {
		fmt.Println("Error opening directory:", err)
		return
	}
	defer dir.Close()

	sourceEntries, err := dir.ReadDir(-1)
	if err != nil {
		fmt.Println("error reading directory:")
		return
	}
	// iterate over source entries
	for _, entry := range sourceEntries {
		entryName := filepath.Base(entry.Name())
		fmt.Println("source entry:", entry)

		tempDirPath := filepath.Join(tempDir, entryName)
		fmt.Println("temporary directory path:", tempDirPath)

		finalDirPath := filepath.Join(finalDir, entryName)
		fmt.Println("finalDirPath:", finalDirPath)

		err := createDir(tempDirPath)
		if err != nil {
			fmt.Printf("Error creating %s %v", tempDirPath, err)
			return // skip this file or dir
		}

		err = createDir(finalDirPath)
		if err != nil {
			fmt.Printf("Error creating %s %v", finalDirPath, err)
			return // skip this file or dir
		}

	}

	walkFunc := func(path string, info os.FileInfo, err error) error {
		// check if there was an error accessing the file or dir
		if err != nil {
			fmt.Printf("Error accessing %s %v", path, err)
			return nil // skip this file or dir
		}

		// check if it is a regular file
		if !info.IsDir() {
			subFolder := strings.Split(path, "\\")[1]
			fileName := strings.TrimSuffix(filepath.Base(path), ".json")
			goName := fileName + ".go"
			jsonPathRev := strings.Replace(path, "\\", "/", -1)
			tempGoPath := tempDir + "/" + subFolder + "/" + goName
			finalGoPath := finalDir + "/" + subFolder + "/" + goName

			//Run quicktype
			err := runQuicktype(jsonPathRev, tempGoPath)
			if err != nil {
				fmt.Println(err)
				return err
			}
			// Copy the generated go file to the final directory
			if err := copyFile(tempGoPath, finalGoPath); err != nil {
				return err
			}

		}
		return nil

	}

	// Traverse the directory tree using Walk
	err = filepath.Walk(sourcePath, walkFunc)
	if err != nil {
		fmt.Printf("Error walking the path %s: %v\n", sourcePath, err)
	}

	//// Remove contents in the temp/bo directory
	//err = removeContentsInDirectory(tempDir)
	//if err != nil {
	//	fmt.Println("Error removing contents in temp/bo directory:", err)
	//	return
	//}

}
