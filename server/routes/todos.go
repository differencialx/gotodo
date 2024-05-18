package routes

import (
	"errors"
	"fmt"
	"gotodo/db"
	"gotodo/models"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

func getTodoList(context *gin.Context) {
	var todos []models.Todo
	result := db.DB.Find(&todos)

	if result.Error != nil {
		internalServerError(context, []gin.H{
			{"message": "Internal server error"},
		})
		return
	}

	successResponse(context, todos)
}

func postTodo(context *gin.Context) {
	var err error
	var todo models.Todo
	err = context.ShouldBindJSON(&todo)
	if err != nil {
		fmt.Println(err)
		badRequest(context, []gin.H{
			{"message": "Could not parse the request"},
		})
		return
	}

	v := validator.New()
	err = v.Struct(todo)
	if err != nil {
		unprocessableEntity(context, extractValidationErrors(todo, err))
		return
	}

	result := db.DB.Create(&todo)
	if result.Error != nil {
		internalServerError(context, []gin.H{
			{"message": "Internal server error"},
		})
		return
	}

	successResponse(context, todo)
}

func getTodo(context *gin.Context) {
	var err error

	todoId, err := strconv.ParseUint(context.Param("id"), 10, 64)
	if err != nil {
		badRequest(context, []gin.H{
			{"message": "Invalid id"},
		})
		return
	}

	var todo models.Todo

	result := db.DB.First(&todo, todoId)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		notFound(context, []gin.H{
			{"message": "Record not found"},
		})
		return
	}

	successResponse(context, todo)
}

func putTodo(context *gin.Context) {
	var err error

	todoId, err := strconv.ParseUint(context.Param("id"), 10, 64)
	if err != nil {
		badRequest(context, []gin.H{
			{"message": "Invalid id"},
		})
		return
	}

	var todo models.Todo

	result := db.DB.First(&todo, todoId)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		notFound(context, []gin.H{
			{"message": "Record not found"},
		})
		return
	}

	err = context.ShouldBindJSON(&todo)
	if err != nil {
		badRequest(context, []gin.H{
			{"message": "Invalid request"},
		})
		return
	}

	v := validator.New()
	err = v.Struct(todo)
	if err != nil {
		unprocessableEntity(context, extractValidationErrors(todo, err))
		return
	}

	result = db.DB.Save(&todo)

	if result.Error != nil {
		internalServerError(context, []gin.H{
			{"message": "Could not save todo"},
		})
	}

	successResponse(context, todo)
}

func deleteTodo(context *gin.Context) {
	var err error

	todoId, err := strconv.ParseUint(context.Param("id"), 10, 64)
	if err != nil {
		badRequest(context, []gin.H{
			{"message": "Invalid id"},
		})
		return
	}

	var todo models.Todo

	result := db.DB.First(&todo, todoId)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		notFound(context, []gin.H{
			{"message": "Record not found"},
		})
		return
	}

	result = db.DB.Delete(&todo)

	if result.Error != nil {
		internalServerError(context, []gin.H{
			{"message": "Could not delete todo"},
		})
		return
	}

	successResponse(context, gin.H{
		"message": "Todo was successfully removed.",
	})
}
