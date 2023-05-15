package models

import (
	"database/sql/driver"
	"errors"
	"github.com/alima12/Blog-Go/database"
	"gorm.io/gorm"
)

type Status string

const (
	Draft     Status = "draft"
	Pending   Status = "pending"
	Published Status = "published"
)

func (s *Status) Scan(value interface{}) error {
	*s = Status(value.(string))
	return nil
}

func (s Status) Value() (driver.Value, error) {
	return string(s), nil
}

type Post struct {
	gorm.Model
	Slug     string `gorm:"not null, unique"`
	Title    string `gorm:"not null"`
	UserID   uint   `gorm:"not null"`
	Content  string `gorm:"not null"`
	ImageURL string
	Status   Status `gorm:"not null;check:status IN ('draft', 'pending', 'published')"`
	Views    int64  `gorm:"default:0"`
}

func (model *Post) GetOne(id string) error {
	db := database.GetDB()
	result := db.First(&model, id)
	if result.Error != nil {
		return errors.New("")
	}
	return nil
}

func (model *Post) Delete(id string) error {
	db := database.GetDB()
	result := db.Delete(&Post{}, "id = ?", id)
	if result.RowsAffected == 1 {
		return nil
	} else {
		return errors.New("post does not exists")
	}
}

func (model *Post) Create() {
	db := database.GetDB()
	db.Create(&model)
}
