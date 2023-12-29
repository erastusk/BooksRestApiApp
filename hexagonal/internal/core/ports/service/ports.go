package ports

import (
	"hexagonal/internal/core/domain"
)

type Repository interface {
	GetAllBooks() []domain.Response
	InsertBook(domain.Request)
	GetBookById(int) domain.Response
	DeleteBook(int) error
	CreateTable()
	CheckIfExist(string) bool
}
