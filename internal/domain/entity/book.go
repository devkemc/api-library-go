package entity

type Book struct {
	Id                int64
	Title             string
	Author            []Author
	ISBN              string
	PublishingCompany string
	Year              int
	Synopsis          string
	QuantityPages     int
	Price             float64
	Availability      bool
	Genre             BookGenre
}
