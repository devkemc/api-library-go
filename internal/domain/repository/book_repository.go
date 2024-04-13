package repository

import (
	"github.com/devkemc/api-library-go/internal/domain/entity"
)

type BookRepository interface {
	CreateBook(book entity.Book) (*entity.Book, error)
	FindBookByID(book entity.Book) (*entity.Book, error)
	FindAllBooks() (*[]entity.Book, error)
	UpdateBook(newBook entity.Book, currentBook entity.Book) (*entity.Book, error)
	DeleteBook(book entity.Book) (*entity.Book, error)
}
