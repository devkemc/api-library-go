package books_usecase

import (
	"ebook-with-go/internal/domain/entity"
	"ebook-with-go/internal/domain/repository"
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
