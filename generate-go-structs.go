package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
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
	return nil
}

func processGoFile(filePath string) error {
	// Read the contents of the Go file
	content, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}
	// Modify the package name based on the directory it was generated from
	dir := filepath.Dir(filePath)
	parts := strings.Split(dir, string(os.PathSeparator))

	if len(parts) < 2 {
		return fmt.Errorf("unexpected directory structure: %s", dir)
	}
	packageName := parts[1]

	// Replace the package declaration in the file content
	newContent := bytes.Replace(content, []byte("package main"), []byte("package "+packageName), 1)

	// Write the modified content back to the file
	err = os.WriteFile(filePath, newContent, 0644)
	if err != nil {
		return err
	}

	return nil
}

func getSubDir(path string, subFolders []string) string {
	pathLower := strings.ToLower(path)

	for _, entry := range subFolders {
		pattern := fmt.Sprintf(`\b%s\b`, regexp.QuoteMeta(entry))
		regex := regexp.MustCompile(pattern)
		if strings.Contains(pathLower, entry) && regex.MatchString(pathLower) {
			return entry
		}
	}
	return "sub directory ot found"
}

func getBaseName(filepath string) string {
	parts := strings.Split(filepath, "\\")
	filename := parts[len(parts)-1]
	base := strings.TrimSuffix(filename, ".json")
	parts2 := strings.Split(base, ".")
	if len(parts2) > 1 {
		// Return the last element
		return parts2[len(parts2)-1]
	}

	return base
}

func main() {
	sourcePath := "bo4e_schemas"
	tempDir := "temp"
	finalDir := "bo4e_structs"

	subfolders := []string{"bo", "com", "enum"}

	// create tempDir and finalDir
	for _, folder := range subfolders {
		tempDirPath := filepath.Join(tempDir, folder)
		finalDirPath := filepath.Join(finalDir, folder)

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
		if !info.IsDir() && filepath.Ext(path) == ".json" {
			subDir := getSubDir(path, subfolders)
			fileName := getBaseName(path)
			goName := fileName + ".go"
			jsonPathRev := strings.Replace(path, "\\", "/", -1)
			tempGoPath := tempDir + "/" + subDir + "/" + goName
			finalGoPath := finalDir + "/" + subDir + "/" + goName

			//Run quicktype
			err := runQuicktype(jsonPathRev, tempGoPath)
			if err != nil {
				fmt.Println(err)
				return err
			}
			err = processGoFile(tempGoPath)
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
	err := filepath.Walk(sourcePath, walkFunc)
	if err != nil {
		fmt.Printf("Error walking the path %s: %v\n", sourcePath, err)
	}

	// Remove temp directory and all the subdirectories
	//err = os.RemoveAll(tempDir)
	//if err != nil {
	//	fmt.Println("Error removing  temp directory:", err)
	//	return
	//}

}
