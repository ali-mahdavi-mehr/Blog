package models

import "time"

type Post struct {
	Title     string
	Author    User
	Body      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
