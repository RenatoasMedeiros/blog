package store

import (
	"database/sql"
)

type Storage struct {
	Posts interface {
	}
}

func NewStorage(db *sql.DB) Storage {
	return Storage{}
}
