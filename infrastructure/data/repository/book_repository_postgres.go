package repository

import (
	"ebook-with-go/domain/entity"
)

type BookRepositoryPostgres struct {
}

func NewBookRepositoryPostgres() *BookRepositoryPostgres {
	return &BookRepositoryPostgres{}
}
func (b *BookRepositoryPostgres) CreateBook(book *entity.Book) (*entity.Book, error) {
	return &entity.Book{}, nil
}
func (b *BookRepositoryPostgres) FindBookByID(id int64) (*entity.Book, error) {
	return &entity.Book{Id: 57, Title: "Golang Programming", Author: "John Doe"}, nil
}

func (b *BookRepositoryPostgres) FindAllBooks() (*[]entity.Book, error) {
	var books = []entity.Book{
		{Id: 57, Title: "Golang Programming", Author: "John Doe"},
		{Id: 63, Title: "Docker Container", Author: "kennedy"},
		{Id: 97, Title: "Kubernetes Orchest", Author: "jose"},
	}
	return &books, nil
}
func (b *BookRepositoryPostgres) UpdateBook(book *entity.Book) (*entity.Book, error) {
	return &entity.Book{}, nil
}
func (b *BookRepositoryPostgres) DeleteBook(id int64) (*entity.Book, error) {
	return &entity.Book{}, nil
}
