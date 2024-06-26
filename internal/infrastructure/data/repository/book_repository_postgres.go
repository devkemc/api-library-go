package repository

import (
	"errors"
	"fmt"
	"github.com/devkemc/api-library-go/internal/domain/entity"
	"github.com/devkemc/api-library-go/internal/infrastructure/data"
	"strings"
)

type BookRepositoryPostgres struct {
	conn *data.Connection
}

func NewBookRepositoryPostgres(conn *data.Connection) *BookRepositoryPostgres {
	return &BookRepositoryPostgres{conn: conn}
}
func (b *BookRepositoryPostgres) CreateBook(book entity.Book) (*entity.Book, error) {
	transaction, err := b.conn.Conn.Begin()
	if err != nil {
		return nil, err
	}
	query := "INSERT INTO books (" +
		"bok_title, " +
		"bok_isbn, " +
		"bok_publish_company, " +
		"bok_year, bok_synopsis, " +
		"bok_quantity_pages, " +
		"bok_price, " +
		"bok_availability, " +
		"bok_bgr_id ) " +
		"VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)" +
		"RETURNING bok_id;"
	row := transaction.QueryRow(
		query,
		book.Title,
		book.ISBN,
		book.PublishingCompany,
		book.Year,
		book.Synopsis,
		book.QuantityPages,
		book.Price,
		book.Availability,
		book.Genre.Id)
	var bookId int64
	if errId := row.Scan(&bookId); errId != nil {
		return nil, errId
	}
	book.Id = bookId
	for _, author := range book.Authors {
		query = "INSERT INTO books_authors (boa_bok_id, boa_ath_id) VALUES ($1, $2);"
		if _, err = transaction.Exec(query, book.Id, author.Id); err != nil {
			if errRollBack := transaction.Rollback(); errRollBack != nil {
				return nil, errRollBack
			}
			return nil, err
		}
	}

	if err = transaction.Commit(); err != nil {
		return nil, err
	}
	return &book, nil
}
func (b *BookRepositoryPostgres) FindBookByID(book entity.Book) (*entity.Book, error) {
	query := "SELECT * FROM books WHERE bok_id = $1"
	row, err := b.conn.Conn.Query(query, book.Id)
	if err != nil {
		return nil, err
	}
	defer row.Close()
	if !row.Next() {
		return nil, errors.New("Book not found")
	}
	book.Genre = entity.BookGenre{}
	err = row.Scan(&book.Id, &book.Title, &book.ISBN, &book.PublishingCompany, &book.Year, &book.Synopsis, &book.QuantityPages, &book.Price, &book.Availability, &book.Genre.Id)
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
	var books []entity.Book
	for rows.Next() {
		var book entity.Book
		book.Genre = entity.BookGenre{}
		err = rows.Scan(
			&book.Id,
			&book.Title,
			&book.ISBN,
			&book.PublishingCompany,
			&book.Year,
			&book.Synopsis,
			&book.QuantityPages,
			&book.Price,
			&book.Availability,
			&book.Genre.Id)
		if err != nil {
			return nil, err
		}
		books = append(books, book)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return &books, nil
}

func (b *BookRepositoryPostgres) UpdateBook(book entity.Book, currentBook entity.Book) (*entity.Book, error) {
	query := "UPDATE books SET "
	var columns []string
	var params []interface{}
	if book.Title != "" {
		columns = append(columns, fmt.Sprintf("bok_title = $%d", len(columns)+1))
		params = append(params, book.Title)
	}
	if book.ISBN != "" {
		columns = append(columns, fmt.Sprintf("bok_isbn = $%d", len(columns)+1))
		params = append(params, book.ISBN)
	}
	if book.PublishingCompany != "" {
		columns = append(columns, fmt.Sprintf("bok_publish_company = $%d", len(columns)+1))
		params = append(params, book.PublishingCompany)
	}
	if book.Year != 0 {
		columns = append(columns, fmt.Sprintf("bok_year = $%d", len(columns)+1))
		params = append(params, book.Year)
	}
	if book.Synopsis != "" {
		columns = append(columns, fmt.Sprintf("bok_synopsis = $%d", len(columns)+1))
		params = append(params, book.Synopsis)
	}
	if book.QuantityPages != 0 {
		columns = append(columns, fmt.Sprintf("bok_quantity_pages = $%d", len(columns)+1))
		params = append(params, book.QuantityPages)
	}
	if book.Price != 0 {
		columns = append(columns, fmt.Sprintf("bok_price = $%d", len(columns)+1))
		params = append(params, book.Price)
	}
	if book.Availability != currentBook.Availability {
		columns = append(columns, fmt.Sprintf("bok_availability = $%d", len(columns)+1))
		params = append(params, book.Availability)
	}
	if book.Genre.Id != 0 {
		columns = append(columns, fmt.Sprintf("bok_bgr_id $%d", len(columns)+1))
		params = append(params, book.Genre.Id)
	}
	query += strings.Join(columns, ", ")
	query += fmt.Sprintf(" WHERE bok_id = $%d", len(columns)+1)
	params = append(params, book.Id)

	_, err := b.conn.Conn.Exec(query, params...)
	if err != nil {
		return nil, err
	}
	return &book, nil
}
func (b *BookRepositoryPostgres) DeleteBook(book entity.Book) (*entity.Book, error) {
	query := "DELETE FROM books WHERE bok_id = $1"
	if _, err := b.conn.Conn.Exec(query, book.Id); err != nil {
		return nil, err
	}
	return &book, nil
}
