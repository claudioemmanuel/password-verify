package main

import (
	adapter "github.com/claudioemmanue/go-api/adapter/api"
	"github.com/claudioemmanue/go-api/application/services"
)

func main() {
	validator := &services.PasswordValidator{}
	router := adapter.SetupRoutes(validator)

	router.Run(":8080")
}
