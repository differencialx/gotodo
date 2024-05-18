package routes

import (
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func successResponse(context *gin.Context, body interface{}) {
	context.JSON(http.StatusOK, gin.H{
		"errors": nil,
		"data":   body,
	})
}

func badRequest(context *gin.Context, body []gin.H) {
	context.JSON(http.StatusBadRequest, errorMessage(body))
}

func notFound(context *gin.Context, body []gin.H) {
	context.JSON(http.StatusNotFound, errorMessage(body))
}

func unprocessableEntity(context *gin.Context, body []gin.H) {
	context.JSON(http.StatusUnprocessableEntity, errorMessage(body))
}

func internalServerError(context *gin.Context, body []gin.H) {
	context.JSON(http.StatusInternalServerError, body)
}

func errorMessage(body []gin.H) gin.H {
	return gin.H{
		"data":   nil,
		"errors": body,
	}
}

var ErrorMessages map[string]string = map[string]string{
	"Text.required": "Text is a required field",
	"Text.min":      "Text must be at least 3 characters long",
}

func getJSONTag(f reflect.StructField) string {
	jsonTag := f.Tag.Get("json")
	if jsonTag == "" {
		return f.Name
	}
	return jsonTag
}

func extractValidationErrors(model interface{}, err error) []gin.H {
	validationMessages := []gin.H{}
	modelType := reflect.TypeOf(model)
	for _, err := range err.(validator.ValidationErrors) {
		field, _ := modelType.FieldByName(err.Field())
		jsonTag := getJSONTag(field)
		validationMessages = append(validationMessages, gin.H{
			jsonTag: ErrorMessages[err.Field()+"."+err.Tag()],
		})
	}

	return validationMessages
}
