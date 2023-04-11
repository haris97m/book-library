package web

import (
	"time"
)

type AuthorCreateRequest struct {
	FirstName   string    `json:"first_name" validate:"required"`
	LastName    string    `json:"last_name" validate:"required"`
	DateOfBirth time.Time `json:"date_of_birth" validate:"required"`
	DateOfDeath time.Time `json:"date_of_death" validate:"default=null"`
}

type AuthorUpdateRequest struct {
	ID          int       `json:"id"`
	FirstName   string    `json:"first_name" validate:"required"`
	LastName    string    `json:"last_name" validate:"required"`
	DateOfBirth time.Time `json:"date_of_birth" validate:"required"`
	DateOfDeath time.Time `json:"date_of_death" validate:"default=null"`
}

type AuthorResponse struct {
	ID          uint      `json:"id"`
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	DateOfBirth time.Time `json:"date_of_birth"`
	DateOfDeath time.Time `json:"date_of_death"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
