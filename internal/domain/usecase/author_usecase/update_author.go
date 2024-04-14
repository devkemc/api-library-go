package author_usecase

import (
	"errors"
	"github.com/devkemc/api-library-go/internal/domain/entity"
	"github.com/devkemc/api-library-go/internal/domain/repository"
)

type UpdateAuthor struct {
	repository     repository.AuthorRepository
	findAuthorById *FindAuthorByID
}

func NewUpdateAuthor(findAuthorById *FindAuthorByID, repository repository.AuthorRepository) *UpdateAuthor {
	return &UpdateAuthor{repository: repository, findAuthorById: findAuthorById}
}
func (u *UpdateAuthor) UpdateAuthor(author entity.Author) (*entity.Author, error) {
	currentAuthor, err := u.findAuthorById.Execute(author)
	if err != nil {
		return nil, errors.New("author not found")

	}
	return u.repository.UpdateAuthor(*currentAuthor, author)
}
