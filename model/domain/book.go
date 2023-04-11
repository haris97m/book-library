package domain

import (
	"time"

	"gorm.io/gorm"
)

// STRUCTURE TABLE
type Book struct {
	ID          uint      `gorm:"primaryKey"`
	Title       string    `gorm:"size:100"`
	Summary     string    `gorm:"type:text"`
	Genre       string    `gorm:"size:50"`
	ReleaseDate time.Time `gorm:"type:date"`
	Author      Author
	AuthorID    uint
	Publisher   Publisher
	PublisherID uint
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}
