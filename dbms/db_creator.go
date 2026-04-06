package dbms

import (
	"database/sql"
	"fmt"
	"os"

	_ "modernc.org/sqlite"
)

/*
 * we can use the most use github.com/mattn/go-sqlite3 as sqlite3 driver
 * but this requires cgo to be enabled in go env for cgo to be enabled we
 * need a gcc compiler along with cgo.
 *
 * to have sqlite3 driver with pure go we can use github.com/modernc-org/sqlite
 */

func CreateSQLiteDB() {
	fmt.Println("-- [SQLite3] Creating tables in Go --")
	os.Remove("demo.db")

	db, err := sql.Open("sqlite", "demo.db")
	if err != nil {
		fmt.Printf("Error creating sqlite db: %+v\n", err)
	}

	dbSchema := ReadDBSchema("create_table.sql")
	_, err = db.Exec(dbSchema)
	if err != nil {
		fmt.Printf("Error while executing schema %+v", err)
	}
}
