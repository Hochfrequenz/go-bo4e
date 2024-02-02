//package main
//
//import (
//	"fmt"
//	"io"
//	"os"
//	"os/exec"
//	"path/filepath"
//	"strings"
//)
//
//func runQuicktype(jsonPath, outputDir string) error {
//	cmd := exec.Command("quicktype", "--src", jsonPath, "-o", outputDir, "--src-lang", "schema", "--lang", "go", "--multi-file-output")
//
//	output, err := cmd.CombinedOutput()
//	if err != nil {
//		return fmt.Errorf("error running quicktype: %s\nOutput:\n%s", err, output)
//	}
//	return nil
//}
//
//func ensureDirectoryExists(directoryPath string) error {
//	// Ensure the directory exists
//	if err := os.MkdirAll(directoryPath, os.ModePerm); err != nil {
//		return err
//	}
//	return nil
//}
//
//func convertToRelativePathAndFileName(sourceDirectory, filePath string, info os.FileInfo) (string, string) {
//	relativePath, _ := filepath.Rel(sourceDirectory, filePath)
//	relativePath = strings.TrimPrefix(relativePath, string(filepath.Separator))
//	jsonFileName := strings.TrimSuffix(info.Name(), filepath.Ext(info.Name()))
//	return relativePath, jsonFileName
//}
//
//func copyGoFileToDestination(temporaryPath, finalDestination, relativePath, jsonFileName string) error {
//	// Assuming that the generated Go file has the same name as the JSON file
//	goFileName := jsonFileName + ".go"
//
//	// Calculate paths
//	temporaryFilePath := filepath.Join(temporaryPath, relativePath, goFileName)
//	finalFilePath := filepath.Join(finalDestination, relativePath, goFileName)
//
//	// Ensure the final destination directory exists
//	if err := ensureDirectoryExists(filepath.Dir(finalFilePath)); err != nil {
//		return err
//	}
//
//	// Copy the Go file to the final destination
//	if err := copyFile(temporaryFilePath, finalFilePath); err != nil {
//		return err
//	}
//
//	return nil
//}
//
//func copyFile(src, dst string) error {
//	srcFile, err := os.Open(src)
//	if err != nil {
//		return err
//	}
//	defer srcFile.Close()
//
//	dstFile, err := os.Create(dst)
//	if err != nil {
//		return err
//	}
//	defer dstFile.Close()
//
//	_, err = io.Copy(dstFile, srcFile)
//	if err != nil {
//		return err
//	}
//
//	return nil
//}
//
//func main() {
//	sourceDirectory := "./bo4e_schemas/bo"
//	temporaryDestination := "./temp/bo"
//	finalDestination := "./destination/bo"
//
//	// create the directories if they don't exist
//	if err := ensureDirectoryExists(temporaryDestination); err != nil {
//		fmt.Println("Error creating temporary directory:", err)
//		return
//	}
//	if err := ensureDirectoryExists(finalDestination); err != nil {
//		fmt.Println("Error creating final destination directory:", err)
//		return
//	}
//
//	err := filepath.Walk(sourceDirectory, func(filePath string, info os.FileInfo, err error) error {
//		if err != nil {
//			return err
//		}
//
//		// Process only JSON files
//		if !info.IsDir() && strings.HasSuffix(info.Name(), ".json") {
//			relativePath, jsonFileName := convertToRelativePathAndFileName(sourceDirectory, filePath, info)
//
//			// Run quicktype
//			err := runQuicktype(filePath, filepath.Join(temporaryDestination, relativePath))
//			if err != nil {
//				fmt.Println(err)
//				return nil
//			}
//
//			// Copy the generated Go file to the final destination
//			err = copyGoFileToDestination(temporaryDestination, finalDestination, relativePath, jsonFileName)
//			if err != nil {
//				fmt.Println("Error copying Go file to destination:", err)
//				return nil
//			}
//		}
//
//		return nil
//	})
//
//	if err != nil {
//		fmt.Println("Error:", err)
//		return
//	}
//
//	fmt.Println("Go files generated and copied successfully.")
//}
