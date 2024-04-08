package repository

import (
	"ebook-with-go/internal/domain/entity"
	"ebook-with-go/internal/infrastructure/data"
	"errors"
)

type BookRepositoryPostgres struct {
	conn *data.Connection
}

func NewBookRepositoryPostgres(conn *data.Connection) *BookRepositoryPostgres {
	return &BookRepositoryPostgres{conn: conn}
}
func (b *BookRepositoryPostgres) CreateBook(book entity.Book) (*entity.Book, error) {
	query := "INSERT INTO books (title, author) VALUES ($1, $2)"
	if _, err := b.conn.Conn.Exec(query, book.Title, book.Author); err != nil {
		return nil, err
	}
	return &book, nil
}
func (b *BookRepositoryPostgres) FindBookByID(book entity.Book) (*entity.Book, error) {
	query := "SELECT * FROM books WHERE id = $1"
	row, err := b.conn.Conn.Query(query, book.Id)
	if err != nil {
		return nil, err
	}
	defer row.Close()
	if !row.Next() {
		return nil, errors.New("Book not found")
	}
	err = row.Scan(&book.Id, &book.Title, &book.Author)
	if err != nil {
		return nil, err
	}
	return &book, nil
}

func (b *BookRepositoryPostgres) FindAllBooks() (*[]entity.Book, error) {
	query := "SELECT * FROM books"
	rows, err := b.conn.Conn.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	books := []entity.Book{}
	for rows.Next() {
		var book entity.Book
		err := rows.Scan(&book.Id, &book.Title, &book.Author)
		if err != nil {
			return nil, err
		}
		books = append(books, book)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return &books, nil
}

func (b *BookRepositoryPostgres) UpdateBook(book entity.Book) (*entity.Book, error) {
	query := "UPDATE books SET title = $1, author = $2 WHERE id = $3"
	if _, err := b.conn.Conn.Exec(query, book.Title, book.Author, book.Id); err != nil {
		return nil, err
	}
	return &book, nil
}
func (b *BookRepositoryPostgres) DeleteBook(book entity.Book) (*entity.Book, error) {
	query := "DELETE FROM books WHERE id = $1"
	if _, err := b.conn.Conn.Exec(query, book.Id); err != nil {
		return nil, err
	}
	return &book, nil
}
