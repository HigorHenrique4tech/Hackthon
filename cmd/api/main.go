package main

import (
	"log"
	"net/http"
	
	"github.com/go-chi/chi/v5"
	"github.com/internal/handlers"
	"hackathon-api/internal/middleware"
)

func main() {
	r := chi.NewRouter()

	// Apply middleware
	r.Use(middleware.SecurityHeaders)
	r.Use(middleware.Logger)

	// Define routes
	itemhandler := handlers.NewItemHandler()
	r.mount("/items", itemhandler.Routes())
	
	// Start the server
	log.Println("API running on :8080")
	http.ListenAndServe(":8080", r)
} 
