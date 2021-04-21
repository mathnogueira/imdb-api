package movie

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/mathnogueira/imdb-api/storage/movie"
)

type searchMoviesRequest struct {
	Terms []string `json:"terms"`
}

type searchMoviesResponse struct {
	Movies []movieDTO
}

type movieDTO struct {
	ID       string   `json:"id"`
	Name     string   `json:"name"`
	Director string   `json:"director"`
	Cast     []string `json:"cast"`
}

// SearchMovies uses the tags provided by the user and find all movies that contain all of those tags
func SearchMovies(c echo.Context, movieRepository *movie.Repository) error {
	terms, err := getTermsFromRequest(c.Request())
	if err != nil {
		return err
	}

	movies := movieRepository.Search(terms)

	response := searchMoviesResponse{
		Movies: make([]movieDTO, 0, len(movies)),
	}
	for _, movie := range movies {
		movieDTO := movieDTO{
			ID:       movie.ID,
			Name:     movie.Name,
			Director: movie.Director,
			Cast:     movie.Cast,
		}
		response.Movies = append(response.Movies, movieDTO)
	}

	c.Response().Header().Set("Content-Type", "application/json")
	return c.JSON(http.StatusOK, response)
}

func getTermsFromRequest(httpRequest *http.Request) ([]string, error) {
	requestBodyBytes, err := ioutil.ReadAll(httpRequest.Body)
	if err != nil {
		return nil, fmt.Errorf("Could not read request body: %w", err)
	}

	var request searchMoviesRequest
	err = json.Unmarshal(requestBodyBytes, &request)
	if err != nil {
		return nil, fmt.Errorf("Could not parse request body as JSON: %w", err)
	}

	return request.Terms, nil
}
