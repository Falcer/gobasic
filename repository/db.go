package repository

import (
	"github.com/jmoiron/sqlx"
	// pq lib import
	_ "github.com/lib/pq"
)

// OpenDB ahhhhh
func OpenDB() (*sqlx.DB, error) {
	db, err := sqlx.Connect("postgres", "user=postgres dbname=go_login password=argadev123 sslmode=disable")
	if err != nil {
		return nil, err
	}

	return db, nil
}
