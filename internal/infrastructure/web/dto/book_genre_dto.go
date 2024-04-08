package dto

type CreateBookGenreDto struct {
	Name string `json:"name"`
}

type BookGenreDto struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}
