package models

import (
	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	Title  string
	Author User `gorm:"embedded"`
	Body   string
}
