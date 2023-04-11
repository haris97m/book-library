package domain

import (
	"time"

	"gorm.io/gorm"
)

// STRUCTURE TABLE
type Author struct {
	ID          uint      `gorm:"primaryKey"`
	FirstName   string    `gorm:"size:100;not null"`
	LastName    string    `gorm:"size:100;not null"`
	DateOfBirth time.Time `gorm:"type:date;not null"`
	DateOfDeath time.Time `gorm:"type:date;null"`
	Books       []Book    `gorm:"foreignKey:AuthorID"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}
