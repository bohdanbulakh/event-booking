package main

import (
	"event-booking/database"
	"event-booking/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

func main() {
	database.InitDB()
	server := gin.Default()
	loadEnv()
	port := os.Getenv("PORT")

	server.GET("/events", getEvents)
	server.POST("/events", createEvent)

	server.Run(":" + port)
}

func getEvents(context *gin.Context) {
	events, exception := models.GetAllEvents()
	if exception != nil {
		context.JSON(
			http.StatusInternalServerError,
			gin.H{"message": "InternalServerError"},
		)
		fmt.Println(exception)
	}

	context.JSON(http.StatusOK, gin.H{"events": events})
}

func createEvent(context *gin.Context) {
	var event models.Event
	exception := context.ShouldBindJSON(&event)

	if exception != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "InvalidBodyException"})
	}

	exception = event.Save()
	if exception != nil {
		context.JSON(
			http.StatusInternalServerError,
			gin.H{"message": "InternalServerError"},
		)
		fmt.Println(exception)
	}

	context.JSON(http.StatusCreated, nil)
}

func loadEnv() {
	err := godotenv.Load(".env.development")
	if err != nil {
		log.Fatal("Error loading .env file" +
			"\n" + err.Error())
	}
}
