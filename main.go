package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"github.com/joho/godotenv"

	"block-cipher/router"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	// Initialize router
	r := gin.Default()

	// Add endpoints to router
	router.ConfigureRouter(r)
	
	// Get port from environment variables and start server
	r.Run(getEnvPortOr("8080"))
}

func getEnvPortOr(port string) string {
	// If `PORT` variable in environment exists, return it
	if envPort := os.Getenv("PORT"); envPort != "" {
		return ":" + envPort
	}
	// Otherwise, return the value of `port` variable from function argument
	return ":" + port
}