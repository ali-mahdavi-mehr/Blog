package models

import (
	"errors"
	"github.com/alima12/Blog-Go/database"
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
