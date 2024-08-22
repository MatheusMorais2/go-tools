package templates

import (
	"fmt"
	"go-tools/pkg/utils"
	"log"
	"os"
	"runtime"
)

func WriteMainPostgresRepositoryFile(basePath string) error {
	path := fmt.Sprintf("%s/internal/adapters/storage/sql/main", basePath)
	data := writeMainPostgresRepositoryData()
	err := os.WriteFile(utils.GoFile(path), data, 0600)
	if err != nil {
		_, file, line, _ := runtime.Caller(0)
		log.Fatal(fmt.Printf("Line: %v, File: %s\nError: %+v\n", line, file, err))
	}
	return nil
}

func writeMainPostgresRepositoryData() []byte {
	data := fmt.Sprintln(`package database

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func OpenConnection() (*sql.DB, error) {
	connectionString := "host=localhost post=5432 user=user password=password dbname=db sslmode=disable"
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}
	
	return db, nil
}`)
	return []byte(data)
}