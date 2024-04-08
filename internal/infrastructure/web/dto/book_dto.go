package dto

type BookDTO struct {
	Id                int64  `json:"id"`
	Title             string `json:"title"`
	Author            AuthorDto
	ISBN              string  `json:"isbn"`
	PublishingCompany string  `json:"publishing_company"`
	Year              int     `json:"year"`
	Synopsis          string  `json:"synopsis"`
	QuantityPages     int     `json:"quantity_pages"`
	Price             float64 `json:"price"`
	Availability      bool    `json:"availability"`
	Genre             BookGenreDto
}
