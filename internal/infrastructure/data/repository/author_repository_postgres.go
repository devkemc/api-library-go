package repository

import (
	"fmt"
	"github.com/devkemc/api-library-go/internal/domain/entity"
	"github.com/devkemc/api-library-go/internal/infrastructure/data"
)

type AuthorRepository struct {
	connection *data.Connection
}

func NewAuthorRepository(connection *data.Connection) *AuthorRepository {
	return &AuthorRepository{connection: connection}
}
func (a *AuthorRepository) CreateAuthor(author entity.Author) (*entity.Author, error) {
	query := "INSERT INTO authors (" +
		"ath_name, " +
		"ath_date_of_birth, " +
		"ath_nationality, " +
		"ath_biography, " +
		"ath_gender_identity) " +
		"VALUES (" +
		"$1, " +
		"$2, " +
		"$3, " +
		"$4, " +
		"$5) " +
		"RETURNING ath_id;"
	row := a.connection.Conn.QueryRow(
		query,
		author.Name,
		author.DateOfBirth,
		author.Nationality,
		author.Biography,
		author.GenderIdentity)
	var authorId int64
	if err := row.Scan(&authorId); err != nil {
		return nil, err
	}
	author.Id = authorId
	return &author, nil
}
func (a *AuthorRepository) FindAuthorByID(author entity.Author) (*entity.Author, error) {
	return nil, nil
}
func (a *AuthorRepository) FindAllAuthors() ([]*entity.Author, error) {
	query := "SELECT * FROM authors;"
	rows, err := a.connection.Conn.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var authors []*entity.Author
	for rows.Next() {
		var author entity.Author
		if err = rows.Scan(
			&author.Id,
			&author.Name,
			&author.DateOfBirth,
			&author.Nationality,
			&author.Biography,
			&author.GenderIdentity); err != nil {
			fmt.Print(err)
			return nil, err
		}
		authors = append(authors, &author)
	}
	return authors, nil
}
func (a *AuthorRepository) UpdateAuthor(newAuthor entity.Author, currentAuthor entity.Author) (*entity.Author, error) {
	return nil, nil
}
func (a *AuthorRepository) DeleteAuthor(author entity.Author) (*entity.Author, error) {
	return nil, nil
}
