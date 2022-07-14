package DatabaseMySQL

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func Init(path string) ( *sql.DB, error) {
	db, err := sql.Open("mysql", path)
	if err != nil {
		return nil, err
	}

	return db, nil
}


func Disconnect( db *sql.DB ) error {
	if err := db.Close(); err != nil {
		return err 
	}
	return nil
}
