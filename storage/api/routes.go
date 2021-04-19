package api

import (
	"github.com/labstack/echo/v4"
	"github.com/mathnogueira/imdb-api/storage/api/movie"
)

func (server *Server) setupRoutes() {
	server.echoInstance.POST("/api/movies", func(c echo.Context) error {
		return movie.CreateMovies(c, server.movieRepository)
	})
}
