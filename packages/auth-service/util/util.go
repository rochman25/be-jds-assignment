package util

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type ErrorResponse struct {
	ErrorMessage []string `json:"error_message"`
}

func Index(g *gin.Engine, name string, version string) {
	g.GET("/", func(context *gin.Context) {
		context.JSON(200, struct {
			Name    string `json:"name"`
			Version string `json:"version"`
		}{
			Name:    name,
			Version: version,
		})
	})
}

func ApiErrorResponse(message string) interface{} {
	return ErrorResponse{
		ErrorMessage: []string{
			message,
		},
	}
}

func GenerateCustomMessages(err error) []string {
	var customMessages []string

	// Check if the error is a validation error
	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		for _, fieldError := range validationErrors {
			// Map validation tags to custom error messages
			switch fieldError.Tag() {
			case "required":
				customMessages = append(customMessages, fmt.Sprintf("%s is required", fieldError.Field()))
			case "min":
				customMessages = append(customMessages, fmt.Sprintf("%s must have at least %s characters", fieldError.Field(), fieldError.Param()))
			case "max":
				customMessages = append(customMessages, fmt.Sprintf("%s cannot exceed %s characters", fieldError.Field(), fieldError.Param()))
			default:
				customMessages = append(customMessages, fmt.Sprintf("%s is invalid", fieldError.Field()))
			}
		}
	}

	return customMessages
}
