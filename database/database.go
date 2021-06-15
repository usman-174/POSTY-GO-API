package database

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/usman-174/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDataBase() *gorm.DB {
	fmt.Println("STARTED DATABASE.go")
	env := os.Getenv("ENV")
	if env == "development" {
		fmt.Println("the env is in development")
		err := godotenv.Load()
		if err != nil {
			log.Fatal("ENV LOAD ERROR = ", err.Error())
		}
		fmt.Println(".env file loaded")
	}
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")
	portx := os.Getenv("PORT")

	fmt.Println("port = database=", portx)
	dsn := fmt.Sprintf("host=%s user=%s password=%s port=%s dbname=%s sslmode=require TimeZone=Asia/Shanghai", dbHost, dbUser, dbPass, dbPort, dbName)
	fmt.Println("dsn= ", dsn)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	fmt.Println("connecting to database")
	if err != nil {
		log.Fatal("DATABASE LOAD ERROR occured")
		panic(err.Error())
	}
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Post{})
	db.AutoMigrate(&models.Like{})
	fmt.Println("Connected succesful")
	return db
}
