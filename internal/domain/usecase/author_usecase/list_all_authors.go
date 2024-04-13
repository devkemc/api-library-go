package author_usecase

import (
	"github.com/devkemc/api-library-go/internal/domain/entity"
	"github.com/devkemc/api-library-go/internal/domain/repository"
)

type ListAllAuthors struct {
	authorRepository repository.AuthorRepository
}

func NewListAllAuthors(authorRepository repository.AuthorRepository) *ListAllAuthors {
	return &ListAllAuthors{authorRepository: authorRepository}
}
func (l *ListAllAuthors) Execute() ([]*entity.Author, error) {
	authors, err := l.authorRepository.FindAllAuthors()
	if err != nil {
		return nil, err
	}
	return authors, nil
}
