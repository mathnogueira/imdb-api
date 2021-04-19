package movie

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/mathnogueira/imdb-api/storage/movie"
)

type createMoviesRequest struct {
	Movies []struct {
		ID       string   `json:"id"`
		Name     string   `json:"name"`
		Director string   `json:"director"`
		Cast     []string `json:"cast"`
	} `json:"movies"`
}

// CreateMovies receives a list of movies and index them in a in-memory data structure
// optimized for search.
func CreateMovies(c echo.Context, movieRepository *movie.Repository) error {
	requestBytes, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		return fmt.Errorf("Could not read request body: %w", err)
	}

	var request createMoviesRequest
	err = json.Unmarshal(requestBytes, &request)

	if err != nil {
		return fmt.Errorf("Could not parse request body as JSON content: %w", err)
	}

	validationErrors := validateCreateMoviesRequest(request)
	if len(validationErrors) > 0 {
		return c.JSON(http.StatusBadRequest, errorResponse{Errors: validationErrors})
	}

	saveMoviesIntoStorage(request, movieRepository)
	return c.NoContent(201)
}

func validateCreateMoviesRequest(request createMoviesRequest) []string {
	errors := make([]string, 0)

	for movieIndex, movie := range request.Movies {
		if movie.ID == "" {
			errors = append(errors, fmt.Sprintf("Movie #%d: ID cannot be empty", movieIndex))
		}

		if movie.Name == "" {
			errors = append(errors, fmt.Sprintf("Movie #%d: Name cannot be empty", movieIndex))
		}

		if movie.Director == "" {
			errors = append(errors, fmt.Sprintf("Movie #%d: Director cannot be empty", movieIndex))
		}

		for castIndex, castMember := range movie.Cast {
			if castMember == "" {
				errors = append(errors, fmt.Sprintf("Movie #%d, Cast #%d: Name cannot be empty", movieIndex, castIndex))
			}
		}
	}

	return errors
}

func saveMoviesIntoStorage(request createMoviesRequest, movieRepository *movie.Repository) {
	for _, movieDTO := range request.Movies {
		movie := movie.Movie{
			ID:       movieDTO.ID,
			Name:     movieDTO.Name,
			Director: movieDTO.Director,
			Cast:     movieDTO.Cast,
		}

		movieRepository.Save(movie)
	}

}
