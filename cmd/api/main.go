package main

import (
	"fmt"
	"github.com/devkemc/api-library-go/internal/domain/usecase/book"
	"github.com/devkemc/api-library-go/internal/infrastructure/data"
	"github.com/devkemc/api-library-go/internal/infrastructure/data/repository"
	"github.com/devkemc/api-library-go/internal/infrastructure/web/resource"
	"github.com/devkemc/api-library-go/internal/infrastructure/web/response"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	postgresConn, errConnection := data.ConnectDB()
	if errConnection != nil {
		return
	}
	bookRepository := repository.NewBookRepositoryPostgres(postgresConn)
	findByid := book.NewFindById(bookRepository)
	searchBook := book.NewSearchBook(bookRepository)
	createBook := book.NewCreateBook(bookRepository, searchBook)
	readBook := book.NewFindAll(bookRepository)
	deleteBook := book.NewDeleteBook(bookRepository, findByid)
	updateBook := book.NewUpdateBook(bookRepository, findByid)
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
