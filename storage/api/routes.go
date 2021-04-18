package api

import (
	"github.com/labstack/echo/v4"
	"github.com/mathnogueira/imdb-api/storage/api/movie"
)

func setupRoutes(echoInstance *echo.Echo) {
	echoInstance.POST("/api/movies", movie.CreateMovies)
}
