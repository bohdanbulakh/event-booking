package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

func main() {
	server := gin.Default()
	loadEnv()
	port := os.Getenv("PORT")

	server.GET("/events", getEvents)

	server.Run(":" + port)
}

func getEvents(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"event": "some useful data"})
}

func loadEnv() {
	err := godotenv.Load(".env.development")
	if err != nil {
		log.Fatal("Error loading .env file" +
			"\n" + err.Error())
	}
}
