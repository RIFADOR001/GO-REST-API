package routes

import (
	"example.com/REST-API/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRouts(server *gin.Engine) {

	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent) // The id parameter is dynamic

	// The group gets as parameter the common root. Then it creates a group
	// of routes
	authenticated := server.Group("/")
	// With the Use method we tell it to run the middleware for all of them
	// before the next action
	authenticated.Use(middlewares.Authenticate)
	authenticated.POST("/events", createEvent)
	authenticated.PUT("/events/:id", updateEvent)
	authenticated.DELETE("/events/:id", deleteEvent)

	server.POST("/signup", signup)
	server.POST("/login", login)
}
