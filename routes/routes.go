package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/iiiamnick/EVENT-MANAGEMENT-REST_API/middleware"
)

func RegisterRoutes(server *gin.Engine) {
	server.GET("/events", getEvents) //GET,POST , PUT, PATCH , DELELTE
	server.GET("/events/:id", getEvent)
	auth := server.Group("/")
	auth.Use(middleware.Authenticate)
	auth.POST("/events", createEvents)
	auth.PUT("/events/:id", updateEvent)
	auth.DELETE("/events/:id", deleteEvent)
	auth.DELETE("/events/:id/register", cancelRegistration)
	auth.POST("/events/:id/register", registerForEvent)
	server.POST("/signup", signup)
	server.POST("/login", login)
}
