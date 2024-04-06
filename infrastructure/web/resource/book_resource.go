package resource

import (
	"ebook-with-go/domain/entity"
	"ebook-with-go/domain/usecase/books_usecase"
	"ebook-with-go/infrastructure/web/response"
	"encoding/json"
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
	response   response.Response
}

func NewBookResource(
	createBook *books_usecase.CreateBook,
	findAll *books_usecase.FindAll,
	updateBook *books_usecase.UpdateBook,
	deleteBook *books_usecase.DeleteBook,
	findById *books_usecase.FindById,
	response response.Response,
) BookResource {
	return BookResource{
		createBook: createBook,
		findAll:    findAll,
		updateBook: updateBook,
		deleteBook: deleteBook,
		findById:   findById,
		response:   response,
	}
}

func (r *BookResource) CreateBook(w http.ResponseWriter, req *http.Request) {
	var book entity.Book
	err := json.NewDecoder(req.Body).Decode(&book)
	if err != nil {
		return
	}
	res, err := r.createBook.Execute(book)
	if err != nil {
		return
	}
	r.response.Created(w, res)
}

func (r *BookResource) UpdateBook(w http.ResponseWriter, req *http.Request) {
	var book entity.Book
	err := json.NewDecoder(req.Body).Decode(&book)
	if err != nil {
		return
	}
	idString := mux.Vars(req)["id"]
	id, err := strconv.ParseInt(idString, 10, 64)
	if err != nil {
		return
	}
	book.Id = id
	execute, err := r.updateBook.Execute(book)
	if err != nil {
		return
	}
	r.response.Ok(w, execute)
}

func (r *BookResource) DeleteBook(w http.ResponseWriter, req *http.Request) {
	idString := mux.Vars(req)["id"]
	id, err := strconv.ParseInt(idString, 10, 64)
	book := entity.Book{
		id,
		"",
		"",
	}
	foundBook, err := r.deleteBook.Execute(book)
	if err != nil {
		return
	}
	r.response.Ok(w, foundBook)
}

func (r *BookResource) ListAllBook(w http.ResponseWriter, req *http.Request) {
	response, err := r.findAll.Execute()
	if err != nil {
		return
	}
	r.response.Ok(w, response)
}

func (r *BookResource) FindBookById(w http.ResponseWriter, req *http.Request) {
	idString := mux.Vars(req)["id"]
	id, err := strconv.ParseInt(idString, 10, 64)
	if err != nil {
		return
	}
	book := entity.Book{id, "", ""}
	response, err := r.findById.Execute(book)
	if err != nil {
		r.response.NotFound(w, response)
		return
	}
	r.response.Ok(w, response)
}
