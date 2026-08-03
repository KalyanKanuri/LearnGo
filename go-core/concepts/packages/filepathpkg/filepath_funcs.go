package filepathpkg

import (
	"fmt"
	"io/fs"
	"path/filepath"
)

func FilePathJoin(file string) {
	path := filepath.Join("logs", file)
	fmt.Println("filepath.Join()", path)
}

func FilePathBase(path string) {
	fmt.Println("filepath.Base()", filepath.Base(path))
}

func FilePathExt(path string) {
	fmt.Println("filepath.Ext()", filepath.Ext(path))
}

func FilePathDir(path string) {
	fmt.Println("filepath.Dir()", filepath.Dir(path))
}

func FilePathClean(path string) {
	fmt.Println("filepath.Clean()", filepath.Clean(path))
}

func FilePathWalkDir(path string) {
	err := filepath.WalkDir(path, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if filepath.Ext(d.Name()) == ".go" {
			fmt.Println("filepath.WalkDir()", d.Name())
		}
		return nil
	})
	if err != nil {
		fmt.Println("Error Walking path", err)
	}
}
