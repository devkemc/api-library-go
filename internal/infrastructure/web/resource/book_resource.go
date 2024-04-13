package resource

import (
	"encoding/json"
	"github.com/devkemc/api-library-go/internal/domain/entity"
	"github.com/devkemc/api-library-go/internal/domain/usecase/book_usecase"
	"github.com/devkemc/api-library-go/internal/infrastructure/web/dto/book"
	"github.com/devkemc/api-library-go/pkg/web/response"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type BookResource struct {
	createBook *book_usecase.CreateBook
	findAll    *book_usecase.FindAll
	updateBook *book_usecase.UpdateBook
	deleteBook *book_usecase.DeleteBook
	findById   *book_usecase.FindById
	response   response.Response
}

func NewBookResource(
	createBook *book_usecase.CreateBook,
	findAll *book_usecase.FindAll,
	updateBook *book_usecase.UpdateBook,
	deleteBook *book_usecase.DeleteBook,
	findById *book_usecase.FindById,
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
	var input book.CreateBookDTOInput
	err := json.NewDecoder(req.Body).Decode(&input)
	if err != nil {
		r.response.InvalidParameters(w, err)
		return
	}
	bookEntity := input.ToEntity()
	if err = bookEntity.ValidateCreate(); err != nil {
		r.response.InvalidParameters(w, err)
		return
	}
	_, useCaseErr := r.createBook.Execute(bookEntity)
	if useCaseErr != nil {
		r.response.BadRequest(w, useCaseErr)
		return
	}
	output := book.CreateBookDTOOutputFromEntity(&bookEntity)
	r.response.Created(w, output)
}

func (r *BookResource) UpdateBook(w http.ResponseWriter, req *http.Request) {
	var input book.CreateBookDTOInput
	err := json.NewDecoder(req.Body).Decode(&input)
	if err != nil {
		r.response.InvalidParameters(w, err)
		return
	}
	idString := mux.Vars(req)["id"]
	id, errParseId := strconv.ParseInt(idString, 10, 64)
	if errParseId != nil {
		r.response.InvalidParameters(w, errParseId)
		return
	}
	bookEntity := input.ToEntity()
	bookEntity.Id = id
	execute, errUpdateBook := r.updateBook.Execute(bookEntity)
	if errUpdateBook != nil {
		r.response.BadRequest(w, errUpdateBook)
		return
	}
	r.response.Ok(w, book.CreateBookDTOOutputFromEntity(execute))
}

func (r *BookResource) DeleteBook(w http.ResponseWriter, req *http.Request) {
	idString := mux.Vars(req)["id"]
	id, err := strconv.ParseInt(idString, 10, 64)
	if err != nil {
		r.response.InvalidParameters(w, err)
		return
	}
	bookEntity := entity.Book{}
	bookEntity.Id = id
	_, errDeleteBook := r.deleteBook.Execute(bookEntity)
	if errDeleteBook != nil {
		r.response.BadRequest(w, errDeleteBook)
		return
	}
	r.response.Ok(w, nil)
}

func (r *BookResource) ListAllBook(w http.ResponseWriter, req *http.Request) {
	books, err := r.findAll.Execute()
	if err != nil {
		r.response.BadRequest(w, err)
		return
	}
	r.response.Ok(w, book.ListAllBookDTOFromEntity(*books))
}

func (r *BookResource) FindBookById(w http.ResponseWriter, req *http.Request) {
	idString := mux.Vars(req)["id"]
	id, err := strconv.ParseInt(idString, 10, 64)
	if err != nil {
		return
	}
	bookEntity := entity.Book{}
	bookEntity.Id = id
	result, err := r.findById.Execute(bookEntity)
	if err != nil || result == nil {
		r.response.NotFound(w, result)
		return
	}
	r.response.Ok(w, book.BookDetailsDTOFromEntity(result))
}
