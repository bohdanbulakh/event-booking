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
	"strconv"
)

func main() {
	database.InitDB()
	server := gin.Default()
	loadEnv()
	port := os.Getenv("PORT")

	server.GET("/events", getEvents)
	server.GET("/event/:id", getEvent)
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
		return
	}

	context.JSON(http.StatusOK, gin.H{"events": events})
}

func getEvent(context *gin.Context) {
	id, exception := strconv.ParseInt(context.Param("id"), 10, 64)
	if exception != nil {
		context.JSON(
			http.StatusBadRequest,
			gin.H{"message": "Cannot parse event id"},
		)
		fmt.Println(exception)
		return
	}

	event, exception := models.GetEventById(id)
	if exception != nil {
		context.JSON(
			http.StatusNotFound,
			gin.H{"message": "NotFoundException"},
		)
		fmt.Println(exception)
		return
	}

	context.JSON(http.StatusOK, gin.H{"event": event})
}

func createEvent(context *gin.Context) {
	var event models.Event
	exception := context.ShouldBindJSON(&event)

	if exception != nil {
		context.JSON(
			http.StatusBadRequest,
			gin.H{"message": "InvalidBodyException"},
		)
		fmt.Println(exception)
		return
	}

	exception = event.Save()
	if exception != nil {
		context.JSON(
			http.StatusInternalServerError,
			gin.H{"message": "InternalServerError"},
		)
		fmt.Println(exception)
		return
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
