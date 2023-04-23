package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

func createPostgresConnection() func() *gorm.DB {
	dataBaseConnection := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Tehran",
		os.Getenv("postgres_host"),
		os.Getenv("postgres_user"),
		os.Getenv("postgres_password"),
		os.Getenv("postgres_db_name"),
		os.Getenv("postgres_port"))
	db, err := gorm.Open(postgres.Open(dataBaseConnection), &gorm.Config{})
	if err != nil {
		panic("error in connection")
	}
	return func() *gorm.DB {
		return db
	}
}

func GetDB() *gorm.DB {
	client := createPostgresConnection()
	return client()
}
