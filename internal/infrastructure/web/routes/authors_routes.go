package routes

import (
	"github.com/devkemc/api-library-go/internal/domain/usecase/author_usecase"
	"github.com/devkemc/api-library-go/internal/infrastructure/data"
	"github.com/devkemc/api-library-go/internal/infrastructure/data/repository"
	"github.com/devkemc/api-library-go/internal/infrastructure/web/handlers"
	"github.com/devkemc/api-library-go/pkg/web/response"
	"github.com/gorilla/mux"
)

func InitAuthorsRoutes(api *mux.Router, connection *data.Connection, response response.Response) {
	authorHandler := initAuthorResource(connection, response)
	routerAuthors := api.PathPrefix("/authors").Subrouter()
	routerAuthors.HandleFunc("/", authorHandler.CreateAuthor).Methods("POST")
	routerAuthors.HandleFunc("/", authorHandler.ListAllAuthor).Methods("GET")
	//routerAuthors.HandleFunc("/{id}", authorHandler.FindAuthorById).Methods("GET")
	//routerAuthors.HandleFunc("/{id}", authorHandler.UpdateAuthor).Methods("PUT")
	//routerAuthors.HandleFunc("/{id}", authorHandler.DeleteAuthor).Methods("DELETE")
}
func initAuthorResource(connection *data.Connection, response response.Response) *handlers.AuthorHandler {
	authorRepository := repository.NewAuthorRepository(connection)
	// findByid := author_usecase.NewFindById(authorRepository)
	// searchAuthor := author_usecase.NewSearchAuthor(authorRepository)
	createAuthor := author_usecase.NewRegisterAuthorInput(authorRepository)
	readAuthor := author_usecase.NewListAllAuthors(authorRepository)
	// deleteAuthor := author_usecase.NewDeleteAuthor(authorRepository, findByid)
	// update
	authorHandler := handlers.NewAuthorHandler(createAuthor, readAuthor, response)
	return authorHandler
}
