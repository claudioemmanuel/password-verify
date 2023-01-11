package routes

import (
	"github.com/claudioemmanue/go-api/pkg/controllers"
	"github.com/gin-gonic/gin"
)

// SetupRouter is responsible for setting up the routes
func SetupRouter(router *gin.Engine) *gin.Engine {

	// Route to verify password strength
	router.POST("/verify", controllers.VerifyPassword)

	return router
}
