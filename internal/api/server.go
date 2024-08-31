package api

import (
	db "github.com/Dennisblay/ordering-app-server/internal/database/sqlc"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// Server serves all HTTP requests
type Server struct {
	router *gin.Engine
	store  db.Store
}

// NewServer creates a new HTTP server and setup routing
func NewServer(store db.Store) (*Server, error) {
	server := &Server{
		router: gin.Default(),
		store:  store,
	}

	// Configure CORS middleware before registering routes
	server.ConfigCORSMiddleWare()

	// Register routes after middleware
	server.RegisterRoutes()

	return server, nil
}

// RegisterRoutes registers all the routes for the server
func (s *Server) RegisterRoutes() {
	// Register user routes
	s.userRoutes()
}

// ConfigCORSMiddleWare configures CORS settings
func (s *Server) ConfigCORSMiddleWare() {
	s.router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173", "https://geocomp.netlify.app/"}, // List the allowed origins
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		ExposeHeaders:    []string{"Content-Disposition"},
		AllowCredentials: true,
	}))
}

// RunServer starts the HTTP server on a specified address
func (s *Server) RunServer(address string) error {
	return s.router.Run(address)
}
