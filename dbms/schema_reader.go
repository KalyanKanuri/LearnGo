package dbms

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func ReadDBSchema(schemaName string) string {
	fmt.Println("-- Reading Schema File --")
	schemaPath := filepath.Join("dbms", "schema", schemaName)
	schemaFile, err := os.Open(schemaPath)
	if err != nil {
		fmt.Printf("Error while reading db schema %+v\n", err)
	}
	bufScanner := bufio.NewScanner(schemaFile)
	var schema strings.Builder
	for bufScanner.Scan() {
		line := bufScanner.Text()
		schema.WriteString(line)
	}
	fmt.Printf("Schema from schema file %s is %+v\n", schemaName, schema.String())
	return schema.String()
}
