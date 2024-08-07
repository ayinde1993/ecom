package db

import (
	"database/sql"

	"github.com/go-sql-driver/mysql"
)

// connection to db
func NewMySQLStorage(cfg mysql.Config) (*sql.DB, error) {
	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		return nil, err
		//	log.Fatal(err)
	}

	return db, nil
}
