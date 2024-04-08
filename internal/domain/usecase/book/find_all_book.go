package book

import (
	"github.com/devkemc/api-library-go/internal/domain/entity"
	"github.com/devkemc/api-library-go/internal/domain/repository"
)

type FindAll struct {
	repository repository.BookRepository
}

func NewFindAll(bookRepository repository.BookRepository) *FindAll {
	return &FindAll{repository: bookRepository}
}
func (r *FindAll) Execute() (*[]entity.Book, error) {
	books, err := r.repository.FindAllBooks()
	if err != nil {
		return nil, err
	}
	return books, nil
}
