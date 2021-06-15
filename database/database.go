package database

import (
	"fmt"
	"log"
	"os"

	"github.com/usman-174/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/joho/godotenv"
)

func ConnectDataBase() *gorm.DB {
	fmt.Println("STARTED DATABASE.go")
	env := os.Getenv("env")
	if env == "development" {

		err := godotenv.Load()
		if err != nil {
			log.Fatal("ENV LOAD ERROR = ", err.Error())
		}
	}
	dsn := os.Getenv("DSN")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("DATABASE LOAD ERROR : ")
		panic(err.Error())
	}
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Post{})
	db.AutoMigrate(&models.Like{})

	return db
}
