package models

import "github.com/alima12/Blog-Go/database"

func MigrateModels() {
	db := database.GetDB()
	err := db.AutoMigrate(&Post{}, &User{})
	if err != nil {
		panic(err)
	}
}
