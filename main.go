package main

import (
	"event-booking/database"
	"event-booking/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	server := gin.Default()

	err := godotenv.Load(".env.development")
	if err != nil {
		log.Fatal("Error loading .env file" +
			"\n" + err.Error())
	}
	port := os.Getenv("PORT")

	database.InitDB()
	routes.RegisterRoutes(server)
	server.Run(":" + port)
}
