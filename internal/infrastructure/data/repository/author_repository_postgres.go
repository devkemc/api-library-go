package repository

import (
	"fmt"
	"github.com/devkemc/api-library-go/internal/domain/entity"
	"github.com/devkemc/api-library-go/internal/infrastructure/data"
	"strings"
)

type AuthorRepositoryPostgres struct {
	connection *data.Connection
}

func NewAuthorRepository(connection *data.Connection) *AuthorRepositoryPostgres {
	return &AuthorRepositoryPostgres{connection: connection}
}
func (a *AuthorRepositoryPostgres) CreateAuthor(author entity.Author) (*entity.Author, error) {
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
func (a *AuthorRepositoryPostgres) FindAuthorByID(author entity.Author) (*entity.Author, error) {
	query := "SELECT * FROM authors WHERE ath_id = $1;"
	row := a.connection.Conn.QueryRow(query, author.Id)
	if err := row.Scan(&author.Id, &author.Name, &author.DateOfBirth, &author.Nationality, &author.Biography, &author.GenderIdentity); err != nil {
		return nil, err
	}
	return &author, nil
}
func (a *AuthorRepositoryPostgres) FindAllAuthors() ([]*entity.Author, error) {
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
			return nil, err
		}
		authors = append(authors, &author)
	}
	return authors, nil
}
func (a *AuthorRepositoryPostgres) UpdateAuthor(newAuthor entity.Author, currentAuthor entity.Author) (*entity.Author, error) {
	query := "UPDATE authors SET "
	var columns []string
	var params []interface{}
	if newAuthor.Name != "" {
		columns = append(columns, fmt.Sprintf("ath_name = $%d", len(columns)+1))
		params = append(params, newAuthor.Name)
	}
	if newAuthor.Nationality != "" {
		columns = append(columns, fmt.Sprintf("ath_nationality = $%d", len(columns)+1))
		params = append(params, newAuthor.Nationality)
	}
	if newAuthor.Biography != "" {
		columns = append(columns, fmt.Sprintf("ath_biography = $%d", len(columns)+1))
		params = append(params, newAuthor.Biography)
	}
	if newAuthor.GenderIdentity != "" {
		columns = append(columns, fmt.Sprintf("ath_gender_identity = $%d", len(columns)+1))
		params = append(params, newAuthor.GenderIdentity)
	}
	if newAuthor.DateOfBirth != currentAuthor.DateOfBirth {
		columns = append(columns, fmt.Sprintf("ath_date_of_birth = $%d", len(columns)+1))
		params = append(params, newAuthor.DateOfBirth)
	}
	query += strings.Join(columns, ", ")
	query += fmt.Sprintf(" WHERE bok_id = $%d", len(columns)+1)
	params = append(params, newAuthor.Id)

	_, err := a.connection.Conn.Exec(query, params...)
	if err != nil {
		return nil, err
	}
	return &newAuthor, nil
}
func (a *AuthorRepositoryPostgres) DeleteAuthor(author entity.Author) (*entity.Author, error) {
	query := "DELETE FROM authors WHERE ath_id = $1;"
	if _, err := a.connection.Conn.Exec(query, author.Id); err != nil {
		return nil, err
	}
	return &author, nil
}
