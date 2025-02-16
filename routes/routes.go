package routes

import (
	"event-booking/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	protected := server.Group("/")
	protected.Use(middleware.Authenticate)

	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent)
	protected.POST("/events", createEvent)
	protected.PATCH("/events/:id", updateEvent)
	protected.DELETE("/events/:id", deleteEvent)

	server.POST("/signup", signup)
	server.POST("/login", login)
}
