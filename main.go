package main

import (
	"event-booking/models"
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
	server.POST("/events", createEvent)

	server.Run(":" + port)
}

func getEvents(context *gin.Context) {
	events := models.GetAllEvents()
	context.JSON(http.StatusOK, gin.H{"events": events})
}

func createEvent(context *gin.Context) {
	var event models.Event
	exception := context.ShouldBindJSON(&event)

	if exception != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "InvalidBodyException"})
	}

	event.Id = 1
	event.Save()
	context.JSON(http.StatusCreated, nil)
}

func loadEnv() {
	err := godotenv.Load(".env.development")
	if err != nil {
		log.Fatal("Error loading .env file" +
			"\n" + err.Error())
	}
}
