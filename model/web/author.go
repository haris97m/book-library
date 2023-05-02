package web

import (
	"encoding/json"
	"time"
)

type myTime time.Time

var _ json.Unmarshaler = &myTime{}

func (mt *myTime) UnmarshalJSON(bs []byte) error {
	var s string
	err := json.Unmarshal(bs, &s)
	if err != nil {
		return err
	}
	t, err := time.ParseInLocation("2006-01-02", s, time.UTC)
	if err != nil {
		return err
	}
	*mt = myTime(t)
	return nil
}

type AuthorCreateRequest struct {
	FirstName   string     `json:"first_name" validate:"required"`
	LastName    string     `json:"last_name" validate:"required"`
	DateOfBirth time.Time  `json:"date_of_birth" validate:"required"`
	DateOfDeath *time.Time `json:"date_of_death" default:"null"`
}

type AuthorUpdateRequest struct {
	ID          int        `json:"id"`
	FirstName   string     `json:"first_name" validate:"required"`
	LastName    string     `json:"last_name" validate:"required"`
	DateOfBirth time.Time  `json:"date_of_birth" validate:"required"`
	DateOfDeath *time.Time `json:"date_of_death" default:"null"`
}

type AuthorResponse struct {
	ID          uint       `json:"id"`
	FirstName   string     `json:"first_name"`
	LastName    string     `json:"last_name"`
	DateOfBirth time.Time  `json:"date_of_birth"`
	DateOfDeath *time.Time `json:"date_of_death"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
}
