package authors_dto

import "github.com/devkemc/api-library-go/internal/domain/entity"

type ListAllAuthorsDtoOutput struct {
	Id             int64  `json:"id"`
	Name           string `json:"name"`
	DateOfBirth    string `json:"date_of_birth"`
	Nationality    string `json:"nationality"`
	Biography      string `json:"biography"`
	GenderIdentity string `json:"gender_identity"`
}

func ListAllAuthorsDtoOutputFromEntity(authors []*entity.Author) []*ListAllAuthorsDtoOutput {
	var authorsDTO []*ListAllAuthorsDtoOutput
	for _, author := range authors {
		authorsDTO = append(authorsDTO, &ListAllAuthorsDtoOutput{
			Id:             author.Id,
			Name:           author.Name,
			DateOfBirth:    author.DateOfBirth.Format("2006-01-02"),
			Nationality:    author.Nationality,
			Biography:      author.Biography,
			GenderIdentity: string(author.GenderIdentity),
		})
	}
	return authorsDTO
}
