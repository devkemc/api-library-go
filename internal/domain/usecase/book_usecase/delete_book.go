package book_usecase

import (
	"github.com/devkemc/api-library-go/internal/domain/entity"
	"github.com/devkemc/api-library-go/internal/domain/repository"
)

type DeleteBook struct {
	repository   repository.BookRepository
	findBookById *FindById
}

func NewDeleteBook(bookRepository repository.BookRepository, findBookById *FindById) *DeleteBook {
	return &DeleteBook{repository: bookRepository, findBookById: findBookById}
}
func (d *DeleteBook) Execute(book entity.Book) (*entity.Book, error) {
	foundBook, err := d.findBookById.Execute(book)
	if err != nil {
		return nil, err
	}
	_, err = d.repository.DeleteBook(book)
	if err != nil {
		return nil, err
	}
	return foundBook, nil
}
