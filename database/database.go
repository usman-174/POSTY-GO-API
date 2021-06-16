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
	if env == "production" {
		fmt.Println("the env is in production")

	} else {
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

	// dsn := fmt.Sprintf("host=localhost user=postgres password=postgres port=5432 dbname=auth sslmode=require")
	dsn := fmt.Sprintf("host=%s user=%s password=%s port=%s dbname=%s sslmode=require TimeZone=Asia/Shanghai", dbHost, dbUser, dbPass, dbPort, dbName)
	fmt.Println("dsn= ", dsn)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	fmt.Println("trying to connect to database")
	if err != nil {
		log.Fatal("DATABASE LOAD ERROR occured")

	}
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Post{})
	db.AutoMigrate(&models.Like{})
	fmt.Println("Connected to database succesfully.")
	return db
}
