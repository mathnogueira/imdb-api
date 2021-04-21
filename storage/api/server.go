package api

import (
	"context"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/mathnogueira/imdb-api/storage/database"
	"github.com/mathnogueira/imdb-api/storage/movie"
	"go.uber.org/zap"
)

// Server represents the HTTP Web Server that will receive the requests for this service
type Server struct {
	Port         int
	echoInstance *echo.Echo

	movieRepository *movie.Repository
	logger          *zap.Logger
}

// NewServer creates a new HTTP Server
func NewServer(port int) *Server {
	memoryStorage := database.NewMemoryStorage()

	logger, err := zap.NewDevelopment(
		zap.Fields(zap.String("app", "storage-api")),
	)

	if err != nil {
		panic(err)
	}

	return &Server{
		Port:            port,
		echoInstance:    echo.New(),
		movieRepository: movie.NewRepository(memoryStorage, logger),
		logger:          logger,
	}
}

// Start the server
func (server *Server) Start() {
	server.setupRoutes()
	server.echoInstance.Use(middleware.Logger())

	portBinding := fmt.Sprintf(":%d", server.Port)

	if err := server.echoInstance.Start(portBinding); err != nil && err != http.ErrServerClosed {
		server.logger.Error("Could not start server", zap.Error(err))
		server.echoInstance.Logger.Fatal(err)
	}
}

// Close makes the server quit gracefully
func (server *Server) Close() error {
	server.logger.Sync()
	return server.echoInstance.Shutdown(context.Background())
}

// GetAddress returns the URL to access the server
func (server *Server) GetAddress() string {
	return fmt.Sprintf("http://localhost:%d", server.Port)
}
