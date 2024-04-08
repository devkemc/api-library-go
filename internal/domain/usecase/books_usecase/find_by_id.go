package books_usecase

import (
	"ebook-with-go/internal/domain/entity"
	"ebook-with-go/internal/domain/repository"
)

type FindById struct {
	repository repository.BookRepository
}

func NewFindById(bookRepository repository.BookRepository) *FindById {
	return &FindById{repository: bookRepository}
}
func (r *FindById) Execute(book entity.Book) (*entity.Book, error) {
	foundBook, err := r.repository.FindBookByID(book)
	if err != nil {
		return nil, err
	}
	return foundBook, nil

}
