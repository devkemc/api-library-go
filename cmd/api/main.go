package main

import (
	"fmt"
	"github.com/devkemc/api-library-go/internal/domain/usecase/book_usecase"
	"github.com/devkemc/api-library-go/internal/infrastructure/data"
	"github.com/devkemc/api-library-go/internal/infrastructure/data/repository"
	"github.com/devkemc/api-library-go/internal/infrastructure/web/resource"
	"github.com/devkemc/api-library-go/pkg/web/response"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	postgresConn, errConnection := data.ConnectDB()
	if errConnection != nil {
		return
	}
	bookRepository := repository.NewBookRepositoryPostgres(postgresConn)
	findByid := book_usecase.NewFindById(bookRepository)
	searchBook := book_usecase.NewSearchBook(bookRepository)
	createBook := book_usecase.NewCreateBook(bookRepository, searchBook)
	readBook := book_usecase.NewFindAll(bookRepository)
	deleteBook := book_usecase.NewDeleteBook(bookRepository, findByid)
	updateBook := book_usecase.NewUpdateBook(bookRepository, findByid)
	responseJson := response.NewJsonResponse()
	bookResource := resource.NewBookResource(createBook, readBook, updateBook, deleteBook, findByid, responseJson)

	r := mux.NewRouter()
	r.HandleFunc("/books/", bookResource.CreateBook).Methods("POST")
	r.HandleFunc("/books/", bookResource.ListAllBook).Methods("GET")
	r.HandleFunc("/books/{id}", bookResource.FindBookById).Methods("GET")
	r.HandleFunc("/books/{id}", bookResource.UpdateBook).Methods("PUT")
	r.HandleFunc("/books/{id}", bookResource.DeleteBook).Methods("DELETE")
	r.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			next.ServeHTTP(w, r)
		})
	})
	err := http.ListenAndServe(":8081", r)
	if err != nil {
		return
	}

	fmt.Print("Server is running on port 8081")
}
