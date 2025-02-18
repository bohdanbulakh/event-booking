package routes

import (
	"event-booking/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func getEvents(context *gin.Context) {
	events, exception := models.GetAllEvents()
	if exception != nil {
		context.AbortWithStatusJSON(
			http.StatusInternalServerError,
			gin.H{"message": "InternalServerError"},
		)
		_ = fmt.Errorf("InternalServerError\n%w", exception)
		return
	}

	context.JSON(http.StatusOK, gin.H{"events": events})
}

func getEvent(context *gin.Context) {
	id, exception := strconv.ParseInt(context.Param("id"), 10, 64)
	if exception != nil {
		context.AbortWithStatusJSON(
			http.StatusBadRequest,
			gin.H{"message": "Cannot parse event id"},
		)
		return
	}

	event, exception := models.GetEventById(id)
	if exception != nil {
		context.AbortWithStatusJSON(
			http.StatusNotFound,
			gin.H{"message": "NotFoundException"},
		)
		return
	}

	context.JSON(http.StatusOK, gin.H{"event": event})
}

func createEvent(context *gin.Context) {
	var event models.Event
	exception := context.ShouldBindJSON(&event)

	if exception != nil {
		context.AbortWithStatusJSON(
			http.StatusBadRequest,
			gin.H{"message": "InvalidBodyException"},
		)
		return
	}

	userId := context.GetInt64("userId")
	if userId == 0 {
		context.AbortWithStatusJSON(
			http.StatusUnauthorized,
			gin.H{"message": "UnauthorizedException"},
		)
		return
	}

	event.UserId = userId
	exception = event.Save()
	if exception != nil {
		context.AbortWithStatusJSON(
			http.StatusInternalServerError,
			gin.H{"message": "InternalServerError"},
		)
		_ = fmt.Errorf("InternalServerError\n%w", exception)
		return
	}

	context.JSON(
		http.StatusCreated,
		gin.H{"event": event},
	)
}

func updateEvent(context *gin.Context) {
	id, exception := strconv.ParseInt(context.Param("id"), 10, 64)
	if exception != nil {
		context.AbortWithStatusJSON(
			http.StatusBadRequest,
			gin.H{"message": "Cannot parse event id"},
		)
		return
	}

	event, exception := models.GetEventById(id)
	if exception != nil {
		context.AbortWithStatusJSON(
			http.StatusNotFound,
			gin.H{"message": "NotFoundException"},
		)
		return
	}

	userId := context.GetInt64("userId")
	if event.UserId != userId {
		context.AbortWithStatusJSON(
			http.StatusUnauthorized,
			gin.H{"message": "UnauthorizedException"},
		)
		return
	}

	var updatedEvent models.Event
	exception = context.ShouldBindJSON(&updatedEvent)

	if exception != nil {
		context.AbortWithStatusJSON(
			http.StatusBadRequest,
			gin.H{"message": "InvalidBodyException"},
		)
		return
	}

	updatedEvent.Id = id
	exception = updatedEvent.Update()
	if exception != nil {
		context.AbortWithStatusJSON(
			http.StatusInternalServerError,
			gin.H{"message": "InternalServerError"},
		)
		_ = fmt.Errorf("InternalServerError\n%w", exception)
		return
	}

	context.JSON(
		http.StatusOK,
		gin.H{"event": updatedEvent},
	)
}

func deleteEvent(context *gin.Context) {
	id, exception := strconv.ParseInt(context.Param("id"), 10, 64)
	if exception != nil {
		context.AbortWithStatusJSON(
			http.StatusBadRequest,
			gin.H{"message": "Cannot parse event id"},
		)
		return
	}

	event, exception := models.GetEventById(id)
	if exception != nil {
		context.AbortWithStatusJSON(
			http.StatusNotFound,
			gin.H{"message": "NotFoundException"},
		)
		return
	}

	userId := context.GetInt64("userId")
	if event.UserId != userId {
		context.AbortWithStatusJSON(
			http.StatusUnauthorized,
			gin.H{"message": "UnauthorizedException"},
		)
		return
	}

	exception = event.Delete()
	if exception != nil {
		context.AbortWithStatusJSON(
			http.StatusInternalServerError,
			gin.H{"message": "InternalServerError"},
		)
		_ = fmt.Errorf("InternalServerError\n%w", exception)
		return
	}

	context.JSON(
		http.StatusOK,
		gin.H{"event": event},
	)
}
