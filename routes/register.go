package routes

import (
	"event-booking/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func registerForEvent(context *gin.Context) {
	userId := context.GetInt64("userId")
	eventId, exception := strconv.ParseInt(
		context.Param("id"), 10, 64)

	if exception != nil {
		context.JSON(
			http.StatusBadRequest,
			gin.H{"message": "Cannot parse event id"},
		)
		return
	}

	event, exception := models.GetEventById(eventId)
	if exception != nil {
		context.JSON(
			http.StatusNotFound,
			gin.H{"message": "NotFoundException"},
		)
		return
	}

	exception = event.Register(userId)
	if exception != nil {
		context.JSON(
			http.StatusInternalServerError,
			gin.H{"message": "InternalServerError"},
		)
		fmt.Println(exception)
		return
	}

	context.JSON(
		http.StatusOK,
		nil,
	)
}

func cancelRegistration(context *gin.Context) {
	userId := context.GetInt64("userId")
	eventId, exception := strconv.ParseInt(
		context.Param("id"), 10, 64)

	if exception != nil {
		context.JSON(
			http.StatusBadRequest,
			gin.H{"message": "Cannot parse event id"},
		)
		return
	}

	event, exception := models.GetEventById(eventId)
	if exception != nil {
		context.JSON(
			http.StatusNotFound,
			gin.H{"message": "NotFoundException"},
		)
		return
	}

	exception = event.CancelRegistration(userId)
	if exception != nil {
		context.JSON(
			http.StatusInternalServerError,
			gin.H{"message": "InternalServerError"},
		)
		fmt.Println(exception)
		return
	}
}
