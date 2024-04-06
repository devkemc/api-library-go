package repository

import "ebook-with-go/domain/entity"

type BookRepository interface {
	CreateBook(book *entity.Book) (*entity.Book, error)
	FindBookByID(id int64) (*entity.Book, error)
	FindAllBooks() (*[]entity.Book, error)
	UpdateBook(book *entity.Book) (*entity.Book, error)
	DeleteBook(id int64) (*entity.Book, error)
}
