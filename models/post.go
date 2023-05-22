package models

import (
	"errors"
	"github.com/alima12/Blog-Go/database"
	"gorm.io/gorm"
	"sync/atomic"
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

func (s Status) Value() (string, error) {
	return string(s), nil
}

type Post struct {
	gorm.Model
	Slug     string `gorm:"not null, unique" json:"slug"`
	Title    string `gorm:"not null" json:"title"`
	UserID   uint   `gorm:"not null" json:"author"`
	Content  string `gorm:"not null" json:"content"`
	ImageURL string `json:"image_url"`
	Status   Status `gorm:"not null;check:status IN ('draft', 'pending', 'published')" json:"status"`
	Views    int64  `gorm:"default:0" json:"view_count"`
}

func (model *Post) GetOne(slug string) error {
	db := database.GetDB()
	result := db.First(&model, "slug = ?", slug)
	if result.Error != nil {
		return errors.New("post not found")
	}
	atomic.AddInt64(&model.Views, 1)
	db.Save(&model)
	return nil
}

func (model *Post) Delete(slug string) error {
	db := database.GetDB()
	result := db.Delete(&Post{}, "slug = ?", slug)
	if result.RowsAffected == 1 {
		return nil
	} else {
		return errors.New("post does not exists")
	}
}

func (model *Post) Create() error {
	db := database.GetDB()
	result := db.Create(&model)
	return result.Error
}
