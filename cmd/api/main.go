package main

import (
	"ebook-with-go/domain/usecase/books_usecase"
	"ebook-with-go/infrastructure/data"
	"ebook-with-go/infrastructure/data/repository"
	"ebook-with-go/infrastructure/web/resource"
	"ebook-with-go/infrastructure/web/response"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	postgresConn, errConnection := data.ConnectDB()
	if errConnection != nil {
		return
	}
	bookRepository := repository.NewBookRepositoryPostgres(postgresConn)
	findByid := books_usecase.NewFindById(bookRepository)
	searchBook := books_usecase.NewSearchBook(bookRepository)
	createBook := books_usecase.NewCreateBook(bookRepository, searchBook)
	readBook := books_usecase.NewFindAll(bookRepository)
	deleteBook := books_usecase.NewDeleteBook(bookRepository, findByid)
	updateBook := books_usecase.NewUpdateBook(bookRepository, findByid)
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
