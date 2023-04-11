package web

import "time"

// STRUCTURE REQUEST
type CreatePublisherRequest struct {
	Name     string `json:"name" validate:"required"`
	Industry string `json:"industry" validate:"required"`
}

type UpdatePublisherRequest struct {
	ID       int    `json:"id" validate:"required"`
	Name     string `json:"name" validate:"required"`
	Industry string `json:"industry" validate:"required"`
}

// STRUCTURE RESPONSE
type PublisherResponse struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	Industry  string    `json:"industry"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
