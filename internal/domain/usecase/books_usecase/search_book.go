package books_usecase

import (
	"ebook-with-go/internal/domain/entity"
	"ebook-with-go/internal/domain/repository"
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
