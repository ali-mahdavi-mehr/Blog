package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	FirstName string
	LastName  string
	Username  string
	Email     *string
}
