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
	env := os.Getenv("ENV")
	if env == "development" {
		fmt.Println("the env is in development")
		err := godotenv.Load()
		if err != nil {
			log.Fatal("ENV LOAD ERROR = ", err.Error())
		}
		fmt.Println(".env file loaded")
	}
	clientUrl := os.Getenv("CLIENT_URL")
	port := os.Getenv("PORT")
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{clientUrl, "clever-montalcini-cedd07.netlify.app"},
		AllowCredentials: true,
	})
	fmt.Println("port===", port)
	handler := c.Handler(server)
	err := http.ListenAndServe(":"+port, handler)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println("SERVER UP AND RUNNING")
	fmt.Println("main.go STOP")
}
