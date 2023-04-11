package helper

import (
	"github.com/haris97m/go-fiber/model/domain"
	"github.com/haris97m/go-fiber/model/web"
)

func ToPublisherResponse(publisher domain.Publisher) web.PublisherResponse {
	return web.PublisherResponse{
		ID:        publisher.ID,
		Name:      publisher.Name,
		Industry:  publisher.Industry,
		CreatedAt: publisher.CreatedAt,
		UpdatedAt: publisher.UpdatedAt,
	}
}

func ToPublisherResponses(publishers []domain.Publisher) (responses []web.PublisherResponse) {
	for _, publisher := range publishers {
		responses = append(responses, ToPublisherResponse(publisher))
	}
	return responses
}
