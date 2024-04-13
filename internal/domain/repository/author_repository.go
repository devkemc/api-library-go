package repository

import "github.com/devkemc/api-library-go/internal/domain/entity"

type AuthorRepository interface {
	CreateAuthor(author entity.Author) (*entity.Author, error)
	FindAuthorByID(author entity.Author) (*entity.Author, error)
	FindAllAuthors() ([]*entity.Author, error)
	UpdateAuthor(newAuthor entity.Author, currentAuthor entity.Author) (*entity.Author, error)
	DeleteAuthor(author entity.Author) (*entity.Author, error)
}
