package DatabaseSQLServer

import (
	"database/sql"
	mssql "github.com/denisenkom/go-mssqldb"
)




func Init(url string) (*sql.DB, error) {
	// Create a new connector object by calling NewConnector
	connector, err := mssql.NewConnector(url)
	if err != nil {
		return nil, err
	}

	// Use SessionInitSql to set any options that cannot be set with the dsn string
	// With ANSI_NULLS set to ON, compare NULL data with = NULL or <> NULL will return 0 rows
	connector.SessionInitSQL = "SET ANSI_NULLS ON"

	// Pass connector to sql.OpenDB to get a sql.DB object
	return sql.OpenDB(connector), nil
}

func Disconnect( db *sql.DB ) error {
	if err := db.Close(); err != nil {
		return err 
	}
	return nil
}