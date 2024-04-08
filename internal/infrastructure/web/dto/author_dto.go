package dto

type AuthorDto struct {
	Id          int64  `json:"id"`
	Name        string `json:"name"`
	DateOfBirth string `json:"date_of_birth"`
}
