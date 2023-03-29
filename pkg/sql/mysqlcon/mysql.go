package mysqlcon

import "database/sql"

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
