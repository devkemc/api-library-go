package handlers

import (
	"encoding/json"
	"github.com/devkemc/api-library-go/internal/domain/usecase/author_usecase"
	"github.com/devkemc/api-library-go/internal/infrastructure/web/dto/authors_dto"
	"github.com/devkemc/api-library-go/pkg/web/response"
	"net/http"
)

type AuthorHandler struct {
	registerAuthor *author_usecase.RegisterAuthor
	listAllAuthor  *author_usecase.ListAllAuthors
	updateAuthor   *author_usecase.UpdateAuthor
	findAuthorById *author_usecase.FindAuthorByID
	response       response.Response
}

func NewAuthorHandler(
	registerAuthor *author_usecase.RegisterAuthor,
	listAllAuthor *author_usecase.ListAllAuthors,
	response response.Response,
	updateAuthor *author_usecase.UpdateAuthor,
	findById *author_usecase.FindAuthorByID) *AuthorHandler {
	return &AuthorHandler{
		registerAuthor: registerAuthor,
		listAllAuthor:  listAllAuthor,
		response:       response,
		findAuthorById: findById,
		updateAuthor:   updateAuthor}
}
func (a *AuthorHandler) CreateAuthor(w http.ResponseWriter, req *http.Request) {
	var input authors_dto.RegisterAuthorInput
	if err := json.NewDecoder(req.Body).Decode(&input); err != nil {
		a.response.InvalidParameters(w, err)
		return
	}
	authorEntity := input.ToEntity()
	createdAuthor, useCaseErr := a.registerAuthor.Execute(*authorEntity)
	if useCaseErr != nil {
		a.response.BadRequest(w, useCaseErr)
		return
	}
	output := authors_dto.RegisterAuthorOutputFromEntity(*createdAuthor)
	a.response.Created(w, output)
}

func (a *AuthorHandler) UpdateAuthor(w http.ResponseWriter, req *http.Request) {

}

func (a *AuthorHandler) FindAuthorById(w http.ResponseWriter, req *http.Request) {

}
func (a *AuthorHandler) ListAllAuthor(w http.ResponseWriter, req *http.Request) {
	authors, err := a.listAllAuthor.Execute()
	if err != nil {
		a.response.InternalServerError(w, err)
		return
	}
	output := authors_dto.ListAllAuthorsDtoOutputFromEntity(authors)
	a.response.Ok(w, output)
}
