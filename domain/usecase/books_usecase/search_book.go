package books_usecase

import (
	"ebook-with-go/domain/entity"
	"ebook-with-go/domain/repository"
)

type SearchBook struct {
	repository repository.BookRepository
}

func NewSearchBook(bookRepository repository.BookRepository) *SearchBook {
	return &SearchBook{repository: bookRepository}
}
func (r *SearchBook) Execute(book *entity.Book) (*entity.Book, error) {
	return &entity.Book{Id: 57, Title: "Golang Programming", Author: "John Doe"}, nil

}
