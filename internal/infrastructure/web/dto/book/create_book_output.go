package book

import "github.com/devkemc/api-library-go/internal/domain/entity"

type CreateBookDTOOutput struct {
	Title             string  `json:"title"`
	ISBN              string  `json:"isbn"`
	PublishingCompany string  `json:"publishing_company"`
	Year              int     `json:"year"`
	Synopsis          string  `json:"synopsis"`
	QuantityPages     int     `json:"quantity_pages"`
	Price             float64 `json:"price"`
	Availability      bool    `json:"availability"`
}

func CreateBookDTOOutputFromEntity(book entity.Book) CreateBookDTOOutput {
	return CreateBookDTOOutput{
		Title:             book.Title,
		ISBN:              book.ISBN,
		PublishingCompany: book.PublishingCompany,
		Year:              book.Year,
		Synopsis:          book.Synopsis,
		QuantityPages:     book.QuantityPages,
		Price:             book.Price,
		Availability:      book.Availability,
	}
}
