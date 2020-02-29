package repository

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

// OpenDB ahhhhh
func OpenDB() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:root123@tcp(172.17.0.2:3306)/gorest")
	if err != nil {
		return nil, err
	}

	return db, nil
}
