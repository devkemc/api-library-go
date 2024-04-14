package author_usecase

import (
	"github.com/devkemc/api-library-go/internal/domain/entity"
	"github.com/devkemc/api-library-go/internal/domain/repository"
)

type FindAuthorByID struct {
	repository repository.AuthorRepository
}

func NewFindAuthorByID(repository repository.AuthorRepository) *FindAuthorByID {
	return &FindAuthorByID{repository: repository}
}

func (f *FindAuthorByID) Execute(authorEntity entity.Author) (*entity.Author, error) {
	return f.repository.FindAuthorByID(authorEntity)
}
