package adaptors

import (
	"database/sql"
	ports "hexagonal/internal/core/ports/repository"
)

type RepoSitory struct {
	repo ports.Database
}

func NewDBRepo(r ports.Database) *RepoSitory {
	return &RepoSitory{
		repo: r,
	}
}

func (db *RepoSitory) GetDB() *sql.DB {
	return db.repo.GetDB()
}
