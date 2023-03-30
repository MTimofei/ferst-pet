package mysqlcon

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func OpenMySQLDB(addr *string) (*sql.DB, error) {
	db, err := sql.Open("mysql", *addr)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}
