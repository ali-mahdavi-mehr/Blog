package models

import (
	"gorm.io/gorm"
	"time"
)

type Post struct {
	gorm.Model
	Title     string
	UserID    uint
	Body      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
