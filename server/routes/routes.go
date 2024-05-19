package routes

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	server.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},            // Allow specific origins
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},     // Allow specific HTTP methods
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept"}, // Allow specific headers
		ExposeHeaders:    []string{"Content-Length"},                   // Expose specific headers
		AllowCredentials: true,                                         // Allow credentials
		MaxAge:           12 * time.Hour,                               // Max age for the preflight request
	}))
	api := server.Group("/api")
	api.GET("/todos", getTodoList)
	api.POST("/todos", postTodo)
	api.GET("/todos/:id", getTodo)
	api.PUT("/todos/:id", putTodo)
	api.DELETE("/todos/:id", deleteTodo)
}
