package resource

import (
	"ebook-with-go/domain/entity"
	"ebook-with-go/domain/usecase/books_usecase"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type BookResource struct {
	createBook *books_usecase.CreateBook
	findAll    *books_usecase.FindAll
	updateBook *books_usecase.UpdateBook
	deleteBook *books_usecase.DeleteBook
	findById   *books_usecase.FindById
}

func NewBookResource(
	createBook *books_usecase.CreateBook,
	findAll *books_usecase.FindAll,
	updateBook *books_usecase.UpdateBook,
	deleteBook *books_usecase.DeleteBook,
	findById *books_usecase.FindById,
) BookResource {
	return BookResource{
		createBook: createBook,
		findAll:    findAll,
		updateBook: updateBook,
		deleteBook: deleteBook,
		findById:   findById,
	}
}

func (r *BookResource) CreateBook(w http.ResponseWriter, req *http.Request) {
	var book entity.Book
	err := json.NewDecoder(req.Body).Decode(&book)
	if err != nil {
		return
	}
	res, err := r.createBook.Execute(&book)
	if err != nil {
		return
	}
	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		return
	}
}

func (r *BookResource) UpdateBook(w http.ResponseWriter, req *http.Request) {
	var book entity.Book
	err := json.NewDecoder(req.Body).Decode(&book)
	execute, err := r.updateBook.Execute(&book)
	if err != nil {
		return
	}
	json.NewEncoder(w).Encode(execute)
}

func (r *BookResource) DeleteBook(w http.ResponseWriter, req *http.Request) {
	idString := mux.Vars(req)["id"]
	fmt.Println(idString)
	id, err := strconv.ParseInt(idString, 10, 64)
	if err != nil {
		return
	}
	book, err := r.findById.Execute(id)
	if err != nil {
		return
	}
	foundBook, err := r.deleteBook.Execute(book)
	if err != nil {
		return
	}
	json.NewEncoder(w).Encode(foundBook)
}

func (r *BookResource) ListAllBook(w http.ResponseWriter, req *http.Request) {
	response, err := r.findAll.Execute()
	if err != nil {
		return
	}
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		return
	}
}

func (r *BookResource) FindBookById(w http.ResponseWriter, req *http.Request) {
	idString := mux.Vars(req)["id"]
	fmt.Println(idString)
	id, err := strconv.ParseInt(idString, 10, 64)
	if err != nil {
		return
	}
	response, err := r.findById.Execute(id)
	if err != nil {
		return
	}
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		return
	}
}
