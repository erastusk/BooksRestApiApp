package adapters

import (
	"database/sql"
	ports "hexagonal/internal/core/ports/service"
)

type BooksApi struct {
	db   *sql.DB
	repo ports.Repository
}

func NewService(db *sql.DB, repo ports.Repository) *BooksApi {
	return &BooksApi{
		db:   db,
		repo: repo,
	}
}
