package main

import (
	"foodieMatch/db"
	"foodieMatch/firebase"
	"foodieMatch/routes"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/rs/cors"
)

func main() {

	log.Println("Starting server...")

	log.Println("Connecting with firebase...")
	if os.Getenv("SKIP_FIREBASE") != "true" {
        if err := firebase.InitFirebase(); err != nil {
            log.Fatalf("Firebase initialization failed: %v", err)
        }
    } else {
        log.Println("Skipping Firebase initialization")
    }
	log.Println("Connecting with database...")
    if os.Getenv("SKIP_DB") != "true" {
	    if err := db.InitDB(); err != nil {
	        log.Fatalf("Database initialization failed: %v", err)
	    }
    } else {
        log.Println("Skipping Database initialization")
    }
	r := routes.SetupRoutes()
	r.GET("/healthz", func(c *gin.Context) {
		log.Println("Health check received")
		c.JSON(200, gin.H{"status": "ok"})
	})

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	})
	handler := c.Handler(r)

	log.Println("Starting server on :8080")
	if err := http.ListenAndServe("0.0.0.0:8080", handler); err != nil {
		log.Fatalf("could not start the server: %v", err)
	}
}
