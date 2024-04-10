package routes

import (
	"github.com/gin-gonic/gin"
)

func RegisterRouts(server *gin.Engine) {

	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent) // The id parameter is dynamic
	server.POST("/events", createEvent)
	server.PUT("/events/:id", updateEvent)
}
