package DatabaseSQLite3

import (
	"os"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)


func Init(path ...string) ( *sql.DB, error) {
	var DefaultPath string = "./data/data.db"
	if len(path) != 0 {
		DefaultPath = path[0]
	}

	db, err := sql.Open("sqlite3", DefaultPath)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func Create(path ...string) error {
	var DefaultPath string = "./data"
	if len(path) != 0 {
		DefaultPath = path[0]
	}

	if err := os.MkdirAll(DefaultPath, 0755); err != nil {
		return err
	}
	if _, err := os.Create(DefaultPath+"/data.db"); err != nil {
		return err
	}

	return nil
}

func Disconnect( db *sql.DB ) error {
	if err := db.Close(); err != nil {
		return err 
	}
	return nil
}