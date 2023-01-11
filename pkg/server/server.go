package server

import (
	"log"

	"github.com/claudioemmanue/go-api/pkg/routes"
	"github.com/gin-gonic/gin"
)

// Server is a struct that contains the port and the gin server
type Server struct {
	port   string
	server *gin.Engine
}

// NewServer creates a new server
func NewServer(port string) *Server {
	return &Server{
		port:   port,
		server: gin.Default(),
	}
}

// Start starts the server
func (s *Server) Start() error {

	// Setup the routes
	router := routes.SetupRouter(s.server)

	// Start the server
	log.Println("Server started on port", s.port)
	return router.Run(":" + s.port)
}
