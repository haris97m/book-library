package domain

import (
	"time"

	"gorm.io/gorm"
)

// STRUCTURE TABLE
type Publisher struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"size:100"`
	Industry  string `gorm:"size:100"`
	Books     []Book `gorm:"foreignKey:PublisherID"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
