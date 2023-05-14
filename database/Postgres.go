package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

var client func() *gorm.DB

func createPostgresConnection() func() *gorm.DB {
	dataBaseConnection := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Tehran",
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB_NAME"),
		os.Getenv("POSTGRES_PORT"))
	db, err := gorm.Open(postgres.Open(dataBaseConnection), &gorm.Config{})
	if err != nil {
		panic("error in connection")
	}
	return func() *gorm.DB {
		return db
	}
}

func GetDB() *gorm.DB {
	if client == nil {
		client = createPostgresConnection()
	}
	return client()
}
