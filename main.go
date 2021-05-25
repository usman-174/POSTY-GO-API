package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/usman-174/app"
)

func main() {
	server := app.Router()
	log.Fatal(http.ListenAndServe(":3000", server))
	fmt.Println("SERVER UP AND RUNNING")

}
