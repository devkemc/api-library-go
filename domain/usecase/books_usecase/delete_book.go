package books_usecase

import (
	"ebook-with-go/domain/entity"
	"ebook-with-go/domain/repository"
)

type DeleteBook struct {
	repository   repository.BookRepository
	findBookById *FindById
}

func NewDeleteBook(bookRepository repository.BookRepository, findBookById *FindById) *DeleteBook {
	return &DeleteBook{repository: bookRepository, findBookById: findBookById}
}
func (d *DeleteBook) Execute(book *entity.Book) (*entity.Book, error) {
	_, err := d.findBookById.Execute(book.Id)
	if err != nil {
		return nil, nil
	}
	_, err = d.repository.DeleteBook(book.Id)
	if err != nil {
		return nil, nil
	}
	return book, nil
}
