package main

import (
	"log"
	"net/http"

	"project-app-portfolio-golang-rafli/internal/router"
)

func main() {
	r := router.NewRouter()

	log.Println("Server running on :8080")
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatal(err)
	}
}
