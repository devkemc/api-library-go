package book

import "github.com/devkemc/api-library-go/internal/domain/entity"

type BookDetailsDTO struct {
	ID                int64   `json:"id"`
	Title             string  `json:"title"`
	ISBN              string  `json:"isbn"`
	PublishingCompany string  `json:"publishing_company"`
	Year              int     `json:"year"`
	Synopsis          string  `json:"synopsis"`
	QuantityPages     int     `json:"quantity_pages"`
	Price             float64 `json:"price"`
	Availability      bool    `json:"availability"`
}

func BookDetailsDTOFromEntity(book *entity.Book) *BookDetailsDTO {
	return &BookDetailsDTO{
		book.Id,
		book.Title,
		book.ISBN,
		book.PublishingCompany,
		book.Year,
		book.Synopsis,
		book.QuantityPages,
		book.Price,
		book.Availability,
	}
}
