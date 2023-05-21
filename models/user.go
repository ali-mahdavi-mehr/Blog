package models

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	FirstName string
	LastName  string
	Username  *string `gorm:"uniqueIndex:idx_username"`
	Email     *string `gorm:"uniqueIndex:idx_email"`
	Password  string
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
	Post      []Post
}
