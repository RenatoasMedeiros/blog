package store

import (
	"context"
	"database/sql"
	"errors"
)

type Post struct {
	ID      int64  `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`

	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	IsDeleted bool   `json:"is_deleted"`
}

type PostStore struct {
	db *sql.DB
}

func (ps *PostStore) Create(ctx context.Context, post *Post) error {
	query := `
	INSERT INTO posts (title, content)
		VALUES ($1, $2)
		RETURNING id, created_at, updated_at
	`
	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()
	err := ps.db.QueryRowContext(ctx, query, post.Title, post.Content).Scan(&post.ID, &post.CreatedAt, &post.UpdatedAt)
	if err != nil {
		return err
	}
	return nil
}

func (ps *PostStore) GetAll(ctx context.Context) ([]*Post, error) {
	query := `SELECT id, title, content, created_at, updated_at FROM posts ORDER BY id ASC WHERE is_deleted = FALSE`
	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	rows, err := ps.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	posts := []*Post{}
	for rows.Next() {
		var post Post
		err := rows.Scan(
			&post.ID,
			&post.Title,
			&post.Content,
			&post.CreatedAt,
			&post.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		posts = append(posts, &post)
	}

	return posts, nil
}

func (ps *PostStore) GetByID(ctx context.Context, id int64) (*Post, error) {
	query := `
		SELECT 
			id,
			title,
			content,
			created_at,
			updated_at,
			is_deleted,
		FROM posts WHERE id = $1
	`
	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()
	var post Post
	err := ps.db.QueryRowContext(ctx, query, id).Scan(
		&post.ID,
		&post.Title,
		&post.Content,
		&post.CreatedAt,
		&post.UpdatedAt,
		&post.IsDeleted,
	)
	if err != nil {
		switch {
		//we can padronise the errors in go like that "error.Is"
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrNotFound

		default:
			return nil, err
		}
	}
	return &post, nil
}

func (ps *PostStore) Delete(ctx context.Context, id int64) error {
	query := `
		UPDATE posts SET is_deleted = TRUE WHERE id = $1
	`
	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()
	_, err := ps.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}
	return nil
}

func (ps *PostStore) Update(ctx context.Context, post *Post) error {

	query := `
		UPDATE posts SET title = $1, content = $2 WHERE id = $3
	`
	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()
	_, err := ps.db.ExecContext(ctx, query, post.Title, post.Content, post.ID)
	if err != nil {
		return err
	}
	return nil
}
