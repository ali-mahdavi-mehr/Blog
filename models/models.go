package models

import "github.com/alima12/Blog-Go/database"

func MigrateModels() {
	db := database.GetDB()
	err := db.AutoMigrate(&User{}, &Post{})
	db.Migrator()
	if err != nil {
		panic(err)
	}
}
