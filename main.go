package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/rs/cors"
	"github.com/usman-174/app"
)

func main() {
	var c *cors.Cors
	server := app.Router()
	env := os.Getenv("ENV")
	if env == "production" {
		fmt.Println("the env is in production")

		client_url := os.Getenv("CLIENT_URL")

		c = cors.New(cors.Options{
			AllowedOrigins:   []string{client_url},
			AllowCredentials: true,
		})

		fmt.Println(".env file loaded")
	} else {
		fmt.Println("the env is in development")
		err := godotenv.Load()
		if err != nil {
			log.Fatal("ENV LOAD ERROR = ", err.Error())
		}
		c = cors.New(cors.Options{
			AllowedOrigins:   []string{"http://localhost:3000"},
			AllowCredentials: true,
		})

	}

	port := os.Getenv("PORT")

	handler := c.Handler(server)
	err := http.ListenAndServe(":"+port, handler)
	if err != nil {
		log.Fatal(err.Error())
	} else {
		fmt.Println("SERVER UP AND RUNNING")
		fmt.Println("main.go STOP")
	}

}
