package controllers

import (
	"github.com/claudioemmanue/go-api/pkg/service"
	"github.com/claudioemmanue/go-api/pkg/types"
	"github.com/gin-gonic/gin"
)

// VerifyPassword is responsible for validating the password
func VerifyPassword(c *gin.Context) {
	var payload types.Payload

	// Check if the request is valid with the struct and bind it to the request variable
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(422, gin.H{"message": err.Error()})
		return
	}

	// Call the service to validate the password
	response, err := service.ValidatePassword(payload)

	// If the password is not valid, return the error
	if err != nil {
		c.JSON(422, gin.H{"message": err.Error()})
		return
	}

	// Check if Verify is false and return the error
	if !response.Verify {
		c.JSON(422, response)
		return
	}

	// Return the response
	c.JSON(200, response)
}
