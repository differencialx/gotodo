package routes

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	api := server.Group("/api")
	api.GET("/todos", getTodoList)
	api.POST("/todos", postTodo)
	api.GET("/todos/:id", getTodo)
	api.PUT("/todos/:id", putTodo)
	api.DELETE("/todos/:id", deleteTodo)
}
