package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/sumitst05/patiently/internal/repository"
	"github.com/sumitst05/patiently/internal/router"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	repository.InitDB()

	r := router.SetupRouter()

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	log.Printf("Server running on port %s", port)

	r.Run(":" + port)
}
