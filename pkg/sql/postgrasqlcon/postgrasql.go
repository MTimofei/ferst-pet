package postgrasqlcon

import "database/sql"

func OpenPostgraSQLDB(addr string) (*sql.DB, error) {
	db, err := sql.Open("postgreas", addr)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}
