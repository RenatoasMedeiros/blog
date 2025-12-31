package store

import (
	"context"
	"database/sql"
	"errors"
	"time"
)

var (
	ErrNotFound          = errors.New("Record not found")
	ErrConflict          = errors.New("Record already exists, or is conflicted")
	QueryTimeoutDuration = time.Second * 8
)

type Storage struct {
	Posts interface {
		Create(context.Context, *Post) error
		GetAll(context.Context) ([]*Post, error)
		GetByID(context.Context, int64) (*Post, error)
		Delete(context.Context, int64) error
		Update(context.Context, *Post) error
	}
}

func NewStorage(db *sql.DB) Storage {
	return Storage{
		Posts: &PostStore{db},
	}
}
