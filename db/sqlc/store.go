package db

import "database/sql"

type Store struct {
	*Queries
	db *sql.DB
}