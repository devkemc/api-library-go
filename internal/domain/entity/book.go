package entity

import "errors"

type Book struct {
	Id                int64
	Title             string
	Authors           []Author
	ISBN              string
	PublishingCompany string
	Year              int
	Synopsis          string
	QuantityPages     int
	Price             float64
	Availability      bool
	Genre             BookGenre
}

func (b *Book) ValidateCreate() error {
	if b.Title == "" {
		return errors.New("title is required")
	}
	if len(b.Authors) == 0 {
		return errors.New("author is required")
	}
	if b.ISBN == "" {
		return errors.New("isbn is required")
	}
	if b.PublishingCompany == "" {
		return errors.New("publishing company is required")
	}
	if b.Year == 0 {
		return errors.New("year is required")
	}
	if b.Synopsis == "" {
		return errors.New("synopsis is required")
	}
	if b.QuantityPages == 0 {
		return errors.New("quantity pages is required")
	}
	if b.Price == 0 {
		return errors.New("price is required")
	}
	if b.Genre.Id == 0 {
		return errors.New("genre is required")
	}
	return nil
}
