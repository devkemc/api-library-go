package book_usecase

import (
	"errors"
	"github.com/devkemc/api-library-go/internal/domain/entity"
	"github.com/devkemc/api-library-go/internal/domain/repository"
)

type CreateBook struct {
	repository repository.BookRepository
	searchBook *SearchBook
}

func NewCreateBook(bookRepository repository.BookRepository, searchBook *SearchBook) *CreateBook {
	return &CreateBook{repository: bookRepository, searchBook: searchBook}
}
func (c *CreateBook) Execute(book entity.Book) (*entity.Book, error) {
	existingBook, _ := c.searchBook.Execute(book)
	if existingBook != nil {
		return nil, errors.New("Book already exist")
	}
	newBook, err := c.repository.CreateBook(book)
	if err != nil {
		return nil, err
	}
	return newBook, nil
}
