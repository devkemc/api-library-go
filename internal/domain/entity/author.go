package entity

import "time"

type Author struct {
	Id             int64
	Name           string
	DateOfBirth    time.Time
	Nationality    string
	Biography      string
	GenderIdentity GenderIdentity
	Books          []Book
}
