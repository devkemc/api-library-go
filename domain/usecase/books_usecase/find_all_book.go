package books_usecase

import (
	"ebook-with-go/domain/entity"
	"ebook-with-go/domain/repository"
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
		return nil, nil
	}
	return books, nil
}
