package book

import (
	"github.com/devkemc/api-library-go/internal/domain/entity"
	"time"
)

type CreateBookDTOInput struct {
	Title             string  `json:"title"`
	AuthorIds         []int64 `json:"author_ids"`
	ISBN              string  `json:"isbn"`
	PublishingCompany string  `json:"publishing_company"`
	Year              int     `json:"year"`
	Synopsis          string  `json:"synopsis"`
	QuantityPages     int     `json:"quantity_pages"`
	Price             float64 `json:"price"`
	Availability      bool    `json:"availability"`
	GenreId           int64   `json:"genre_id"`
}

func (c *CreateBookDTOInput) ToEntity() entity.Book {
	var authors []entity.Author
	for _, id := range c.AuthorIds {
		authors = append(authors, entity.Author{id, "", time.Time{}, "", "", "", nil})
	}
	return entity.Book{
		0,
		c.Title,
		authors,
		c.ISBN,
		c.PublishingCompany,
		c.Year,
		c.Synopsis,
		c.QuantityPages,
		c.Price,
		c.Availability,
		entity.BookGenre{c.GenreId, ""},
	}
}
