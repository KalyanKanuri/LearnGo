package main

import (
	"fmt"
	"io/fs"
	"path/filepath"
)

func main() {
	err := filepath.WalkDir("project", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if d.IsDir() && d.Name() == "node_modules" {
			// skip processing node modules
			return filepath.SkipDir
		} else if filepath.Ext(path) == ".log" {
			relPath, err := filepath.Rel("project", path)
			if err != nil {
				return err
			}
			fmt.Printf("Found log file: %s\n", relPath)
		}
		return nil
	})
	if err != nil {
		fmt.Println("Error walking directory", err)
	}
}
