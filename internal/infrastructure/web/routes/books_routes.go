package routes

import (
	"github.com/devkemc/api-library-go/internal/domain/usecase/book_usecase"
	"github.com/devkemc/api-library-go/internal/infrastructure/data"
	"github.com/devkemc/api-library-go/internal/infrastructure/data/repository"
	"github.com/devkemc/api-library-go/internal/infrastructure/web/handlers"
	"github.com/devkemc/api-library-go/pkg/web/response"
	"github.com/gorilla/mux"
)

func InitBooksRoutes(api *mux.Router, connection *data.Connection) {

	bookHandler := initBookResource(connection)
	routerBooks := api.PathPrefix("/books").Subrouter()
	routerBooks.HandleFunc("/", bookHandler.CreateBook).Methods("POST")
	routerBooks.HandleFunc("/", bookHandler.ListAllBook).Methods("GET")
	routerBooks.HandleFunc("/{id}", bookHandler.FindBookById).Methods("GET")
	routerBooks.HandleFunc("/{id}", bookHandler.UpdateBook).Methods("PUT")
	routerBooks.HandleFunc("/{id}", bookHandler.DeleteBook).Methods("DELETE")
}
func initBookResource(postgresConn *data.Connection) *handlers.BookHandler {
	bookRepository := repository.NewBookRepositoryPostgres(postgresConn)
	findByid := book_usecase.NewFindById(bookRepository)
	searchBook := book_usecase.NewSearchBook(bookRepository)
	createBook := book_usecase.NewCreateBook(bookRepository, searchBook)
	readBook := book_usecase.NewFindAll(bookRepository)
	deleteBook := book_usecase.NewDeleteBook(bookRepository, findByid)
	updateBook := book_usecase.NewUpdateBook(bookRepository, findByid)
	responseJson := response.NewJsonResponse()
	bookResource := handlers.NewBookResource(createBook, readBook, updateBook, deleteBook, findByid, responseJson)
	return &bookResource
}
