package book

import "github.com/devkemc/api-library-go/internal/domain/entity"

type ListAllBookDTO struct {
	Id                int64   `json:"id"`
	Title             string  `json:"title"`
	ISBN              string  `json:"isbn"`
	PublishingCompany string  `json:"publishing_company"`
	Year              int     `json:"year"`
	Synopsis          string  `json:"synopsis"`
	QuantityPages     int     `json:"quantity_pages"`
	Price             float64 `json:"price"`
	Availability      bool    `json:"availability"`
}

func ListAllBookDTOFromEntity(books []entity.Book) []ListAllBookDTO {
	var booksDTO []ListAllBookDTO
	for _, book := range books {
		bookDTO := ListAllBookDTO{
			Id:                book.Id,
			Title:             book.Title,
			ISBN:              book.ISBN,
			PublishingCompany: book.PublishingCompany,
			Year:              book.Year,
			Synopsis:          book.Synopsis,
			QuantityPages:     book.QuantityPages,
			Price:             book.Price,
			Availability:      book.Availability,
		}
		booksDTO = append(booksDTO, bookDTO)
	}
	return booksDTO
}
