package book_usecase

import (
	"github.com/devkemc/api-library-go/internal/domain/entity"
	"github.com/devkemc/api-library-go/internal/domain/repository"
)

type UpdateBook struct {
	repository   repository.BookRepository
	findBookById *FindById
}

func NewUpdateBook(bookRepository repository.BookRepository, findBookById *FindById) *UpdateBook {
	return &UpdateBook{repository: bookRepository, findBookById: findBookById}
}
func (r *UpdateBook) Execute(book entity.Book) (*entity.Book, error) {
	_, err := r.findBookById.Execute(book)
	if err != nil {
		return &entity.Book{}, nil
	}
	updatedBook, errUpdateBook := r.repository.UpdateBook(book)
	if errUpdateBook != nil {
		return nil, nil
	}
	return updatedBook, nil
}