package authors_dto

import (
	"github.com/devkemc/api-library-go/internal/domain/entity"
	"time"
)

type RegisterAuthorInput struct {
	Name           string `json:"name"`
	DateOfBirth    string `json:"date_of_birth"`
	Nationality    string `json:"nationality"`
	Biography      string `json:"biography"`
	GenderIdentity string `json:"gender_identity"`
}

func (r *RegisterAuthorInput) ToEntity() *entity.Author {
	dateLayout := "2006-01-02"
	dateOfBirth, _ := time.Parse(dateLayout, r.DateOfBirth)
	return &entity.Author{
		Name:           r.Name,
		DateOfBirth:    dateOfBirth,
		Nationality:    r.Nationality,
		Biography:      r.Biography,
		GenderIdentity: entity.GenderIdentity(r.GenderIdentity),
	}
}
