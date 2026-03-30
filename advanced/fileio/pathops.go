package fileio

import (
	"fmt"
	"os"
	"path/filepath"
)

func HandleFilePaths() {
	fmt.Println("-- Handling File Paths in Go --")
	filePath := filepath.Join("D:\\", "Dev", "Learning")

	os.Chdir(filepath.Clean(filePath))
	dir, _ := os.Getwd()
	fmt.Println("current directory after os.Chdir()", dir)

	os.Chdir("Go")
	absPath, _ := filepath.Abs("go.mod")
	fmt.Println("absolute path", absPath)

	relPath, _ := filepath.Rel("D:\\Dev\\Learning\\Go", "D:\\Dev\\Learning\\Go\\go.mod")
	fmt.Println("relative path", relPath)

	fileExt := filepath.Ext("go.mod")
	fmt.Println("file extension", fileExt)
}
