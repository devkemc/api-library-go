package books_usecase

import (
	"ebook-with-go/domain/entity"
	"ebook-with-go/domain/repository"
	"errors"
)

type SearchBook struct {
	repository repository.BookRepository
}

func NewSearchBook(bookRepository repository.BookRepository) *SearchBook {
	return &SearchBook{repository: bookRepository}
}
func (r *SearchBook) Execute(book entity.Book) (*entity.Book, error) {
	return nil, errors.New("sdds")

}
