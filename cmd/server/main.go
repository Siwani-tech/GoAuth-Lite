package main

import (
	"log"
	"net/http"

	"github.com/Siwani-tech/GoAuth-Lite.git/internal/handlers"
	"github.com/Siwani-tech/GoAuth-Lite.git/internal/middleware"
)

func main() {
	log.Println("Starting GoAuthserver on :8080")
	http.HandleFunc("/health", handlers.HealthHandler)
	http.HandleFunc("/signup", handlers.SignpHandler)
	http.HandleFunc("/login", handlers.LoginHandler)
	http.HandleFunc("/profile", middleware.AuthMiddleware(handlers.ProfileHandler))
	log.Fatal(http.ListenAndServe(":8080", nil))
}

//
