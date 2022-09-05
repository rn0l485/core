package DatabaseMySQL

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

// docker exec benz_mysql_1 /usr/bin/mysqldump -u root --password=2Wr#3z@YUS --all-databases > /home/ubuntu/Benz/backup/$(date +\%Y-\%m-\%d)_all.sql

// sqlcmd -S localhost -U SA -Q "BACKUP DATABASE [demodb] TO DISK = N'/var/opt/mssql/data/demodb.bak' WITH NOFORMAT, NOINIT, NAME = 'demodb-full', SKIP, NOREWIND, NOUNLOAD, STATS = 10"


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
