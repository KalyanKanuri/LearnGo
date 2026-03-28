package fileio

import (
	"embed"
	"fmt"
	"os"
)

func HandlerDirsInGo() {
	fmt.Println("-- Handling Directories in Go --")
	logDir := "logs/app/interface"
	os.Mkdir("logs", 0644) // for creating single directory we can use os.Mkdir
	fmt.Println("Created logs dir using os.Mkdir()")

	// nested dir creation
	os.MkdirAll(logDir, 0644) // if we use os.Mkdir here it will fail
	fmt.Println("Created logs with children using os.MkdirAll()")

	// removing also works same os.Remove() for single directory os.RemoveAll() for nested directories
	os.RemoveAll("logs") // we just need to provide parent dir it will remove it's child dirs as well
	fmt.Println("Removed logs using os.RemoveAll()")

	fileInfo, _ := os.Stat("go.mod")
	fmt.Println(fileInfo.Size(), "bytes")
}

// below is not a comment it's the syntax for embedding static files in Go
//go:embed public
var public embed.FS

func EmbedFilesInGo() {
	fmt.Println("-- Embedding Static Files in Go --")
	data, _ := public.ReadFile("public/text.md")
	fmt.Println(string(data))
}