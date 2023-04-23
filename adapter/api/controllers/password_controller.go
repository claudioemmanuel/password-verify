package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	domain "github.com/claudioemmanue/go-api/domain/entities"
)

func ValidatePasswordEndpoint(validator domain.Validator) gin.HandlerFunc {
	return func(c *gin.Context) {
		var userInput domain.UserInput
		err := c.BindJSON(&userInput)
		if err != nil {
			c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "Invalid request format"})
			return
		}

		isValid, err := validator.ValidatePassword(userInput)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"isValid": isValid})
	}
}