package book

import (
	"ebook-with-go/domain/entity"
	"encoding/json"
	"net/http"
)

type CreateBookDTO struct {
	Title             string  `json:"title"`
	AuthorId          int64   `json:"author_id"`
	ISBN              string  `json:"isbn"`
	PublishingCompany string  `json:"publishing_company"`
	Year              int     `json:"year"`
	Synopsis          string  `json:"synopsis"`
	QuantityPages     int     `json:"quantity_pages"`
	Price             float64 `json:"price"`
	Availability      bool    `json:"availability"`
	GenreId           int64   `json:"genre_id"`
}

func ToEntityFromRequest(req *http.Request) (entity.Book, error) {
	var dto CreateBookDTO
	err := json.NewDecoder(req.Body).Decode(&dto)
	if err != nil {
		return entity.Book{}, err
	}
	return entity.Book{
		Id:                0,
		Title:             dto.Title,
		Author:            nil,
		ISBN:              dto.ISBN,
		PublishingCompany: dto.PublishingCompany,
		Year:              dto.Year,
		Synopsis:          dto.Synopsis,
		QuantityPages:     dto.QuantityPages,
		Price:             dto.Price,
		Availability:      dto.Availability,
		Genre:             entity.BookGenre{Id: dto.GenreId},
	}, nil
}
