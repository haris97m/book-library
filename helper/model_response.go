package helper

import (
	"github.com/haris97m/go-fiber/model/domain"
	"github.com/haris97m/go-fiber/model/web"
)

func ToAuthorResponse(author domain.Author) web.AuthorResponse {
	return web.AuthorResponse{
		ID:          author.ID,
		FirstName:   author.FirstName,
		LastName:    author.LastName,
		DateOfBirth: author.DateOfBirth,
		DateOfDeath: author.DateOfDeath,
		CreatedAt:   author.CreatedAt,
		UpdatedAt:   author.CreatedAt,
	}
}

func ToAuthorResponses(authors []domain.Author) (responses []web.AuthorResponse) {
	for _, author := range authors {
		responses = append(responses, ToAuthorResponse(author))
	}
	return responses
}
