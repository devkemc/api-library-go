package books_usecase

import (
	"ebook-with-go/internal/domain/entity"
	"ebook-with-go/internal/domain/repository"
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
		return nil, nil
	}
	_, err = d.repository.DeleteBook(book)
	if err != nil {
		return nil, nil
	}
	return foundBook, nil
}
