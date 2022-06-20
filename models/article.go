package models

import (
	"time"

	"gorm.io/gorm"
)

type Articles struct {
	ID        uint64          `json:"id"`
	Title     string          `json:"title"`
	Author    string          `json:"author"`
	Genre     string          `json:"genre"`
	ImageUrl  string          `json:"image_url"`
	Created   time.Time       `json:"created"`
	Price     float64         `json:"price"`
	Body      string          `json:"body"`
	CreatedAt time.Time       `json:"-"`
	UpdatedAt time.Time       `json:"-"`
	DeletedAt *gorm.DeletedAt `json:"-"`
}
