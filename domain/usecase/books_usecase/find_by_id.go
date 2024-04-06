package books_usecase

import (
	"ebook-with-go/domain/entity"
	"ebook-with-go/domain/repository"
)

type FindById struct {
	repository repository.BookRepository
}

func NewFindById(bookRepository repository.BookRepository) *FindById {
	return &FindById{repository: bookRepository}
}
func (r *FindById) Execute(id int64) (*entity.Book, error) {
	book, err := r.repository.FindBookByID(id)
	if err != nil {
		return nil, nil
	}
	return book, nil

}
