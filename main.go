package main

import (
	"log"
	"net/http"
	"github.com/rs/cors"
	"foodieMatch/routes"
)


func main() {
    
    c := cors.New(cors.Options{
        AllowedOrigins:   []string{"*"}, // Allow all origins
        AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
        AllowedHeaders:   []string{"Content-Type", "Authorization"},
        AllowCredentials: true,
    })

    r := routes.SetupRoutes()
    handler := c.Handler(r)


    log.Println("Starting server on :8080")
    if err := http.ListenAndServe("0.0.0.0:8080", handler); err!= nil {
        log.Fatalf("could not start the server: %v", err)
    } 
}
