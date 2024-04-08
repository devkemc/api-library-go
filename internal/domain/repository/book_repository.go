package repository

import (
	"github.com/devkemc/api-library-go/internal/domain/entity"
)

type BookRepository interface {
	CreateBook(book entity.Book) (*entity.Book, error)
	FindBookByID(book entity.Book) (*entity.Book, error)
	FindAllBooks() (*[]entity.Book, error)
	UpdateBook(book entity.Book) (*entity.Book, error)
	DeleteBook(book entity.Book) (*entity.Book, error)
}
