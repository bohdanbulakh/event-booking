package routes

import (
	"event-booking/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	protected := server.Group("/")
	protected.Use(middleware.Authenticate)

	server.POST("/signup", signup)
	server.POST("/login", login)

	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent)

	protected.POST("/events", createEvent)
	protected.PATCH("/events/:id", updateEvent)
	protected.DELETE("/events/:id", deleteEvent)
	protected.POST("/events/:id/register", registerForEvent)
	protected.DELETE("/events/:id/register", cancelRegistration)
}
