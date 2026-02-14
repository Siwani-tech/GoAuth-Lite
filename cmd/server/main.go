package main

import (
	"log"
	"net/http"

	"github.com/Siwani-tech/GoAuth-Lite.git/internal/handlers"
)

func main() {
	log.Println("Starting GoAuthserver on :8080")
	http.HandleFunc("/health", handlers.HealthHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
