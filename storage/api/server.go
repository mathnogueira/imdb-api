package api

import (
	"fmt"

	"github.com/labstack/echo/v4"
)

// Server represents the HTTP Web Server that will receive the requests for this service
type Server struct {
	Port         int
	echoInstance *echo.Echo
}

// NewServer creates a new HTTP Server
func NewServer(port int) *Server {
	return &Server{
		Port:         port,
		echoInstance: echo.New(),
	}
}

// Start the server
func (server *Server) Start() {
	setupRoutes(server.echoInstance)

	portBinding := fmt.Sprintf(":%d", server.Port)
	server.echoInstance.Logger.Fatal(server.echoInstance.Start(portBinding))
}

// Close makes the server quit gracefully
func (server *Server) Close() error {
	return server.echoInstance.Close()
}

// GetAddress returns the URL to access the server
func (server *Server) GetAddress() string {
	return fmt.Sprintf("http://localhost:%d", server.Port)
}
