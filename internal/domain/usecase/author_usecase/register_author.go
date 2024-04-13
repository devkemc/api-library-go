package author_usecase

import (
	"errors"
	"github.com/devkemc/api-library-go/internal/domain/entity"
	"github.com/devkemc/api-library-go/internal/domain/repository"
)

type RegisterAuthor struct {
	authorRepository repository.AuthorRepository
}

func NewRegisterAuthorInput(authorRepository repository.AuthorRepository) *RegisterAuthor {
	return &RegisterAuthor{authorRepository: authorRepository}
}

func (r *RegisterAuthor) Execute(author entity.Author) (*entity.Author, error) {
	foundAuthor, _ := r.authorRepository.FindAuthorByID(author)
	if foundAuthor != nil {
		return nil, errors.New("author already exists")
	}
	createdAuthor, err := r.authorRepository.CreateAuthor(author)
	if err != nil {
		return nil, err
	}

	return createdAuthor, nil
}
