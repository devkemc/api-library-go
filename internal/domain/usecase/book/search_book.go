package book

import (
	"errors"
	"github.com/devkemc/api-library-go/internal/domain/entity"
	"github.com/devkemc/api-library-go/internal/domain/repository"
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
