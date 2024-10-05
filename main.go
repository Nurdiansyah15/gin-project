package main

import (
	"gin-project/config"
	"gin-project/routes"
	"gin-project/seeder"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	r := gin.Default()

	// Setup database and other configs
	config.Setup()

	seeder.SeedUsers()

	// Setup routes
	routes.SetupRoutes(r)

	// Run the server
	r.Run(":8080") // Listen on port 8080
}
