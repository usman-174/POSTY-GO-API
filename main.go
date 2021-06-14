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
	fmt.Println("main.go start")
	server := app.Router()
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	clientUrl := os.Getenv("CLIENT_URL")
	port := os.Getenv("PORT")
	fmt.Println("port===", port)
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{clientUrl},
		AllowCredentials: true,
	})

	handler := c.Handler(server)
	err = http.ListenAndServe(port, handler)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println("SERVER UP AND RUNNING")
	fmt.Println("main.go STOP")
}
