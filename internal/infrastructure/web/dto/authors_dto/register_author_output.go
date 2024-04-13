package authors_dto

import "github.com/devkemc/api-library-go/internal/domain/entity"

type RegisterAuthorOutput struct {
	Id             int64  `json:"id"`
	Name           string `json:"name"`
	DateOfBirth    string `json:"date_of_birth"`
	Nationality    string `json:"nationality"`
	Biography      string `json:"biography"`
	GenderIdentity string `json:"gender_identity"`
}

func RegisterAuthorOutputFromEntity(entity entity.Author) *RegisterAuthorOutput {
	return &RegisterAuthorOutput{
		Id:             entity.Id,
		Name:           entity.Name,
		DateOfBirth:    entity.DateOfBirth.String(),
		Nationality:    entity.Nationality,
		Biography:      entity.Biography,
		GenderIdentity: string(entity.GenderIdentity),
	}
}
