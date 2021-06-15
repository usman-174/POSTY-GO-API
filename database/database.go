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
	if env != "" && env == "development" {

		err := godotenv.Load()
		if err != nil {
			log.Fatal("ENV LOAD ERROR = ", err.Error())
		}
	}
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")

	dsn := fmt.Sprintf("host=%s user=%s password=%s port=%s dbname=%s sslmode=require TimeZone=Asia/Shanghai", dbHost, dbUser, dbPass, dbPort, dbName)
	fmt.Println(dsn)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("DATABASE LOAD ERROR occured")
		panic(err.Error())
	}
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Post{})
	db.AutoMigrate(&models.Like{})

	return db
}
