package adapter

import (
	"github.com/gin-gonic/gin"

	"github.com/claudioemmanue/go-api/adapter/api/controllers"
	domain "github.com/claudioemmanue/go-api/domain/entities"
)

func SetupRoutes(validator domain.Validator) *gin.Engine {
	router := gin.Default()

	router.POST("/validate-password", controllers.ValidatePasswordEndpoint(validator))

	return router
}
