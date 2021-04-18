package movie

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/avast/retry-go"
)

// Storage represents the service that is responsible for storing and indexing movies
type Storage struct {
	options StorageOptions
}

// StorageOptions contains all configuration necessary to allow storage to connect to the storage service
type StorageOptions struct {
	StorageURL string
}

// NewStorage creates a new storage instance
func NewStorage(options StorageOptions) *Storage {
	return &Storage{
		options: options,
	}
}

// createMoviesRequest is used to serialize the DTO to JSON
type createMoviesRequest struct {
	Movies []movieDTO `json:"movies"`
}

type movieDTO struct {
	ID       string   `json:"id"`
	Name     string   `json:"name"`
	Director string   `json:"director"`
	Cast     []string `json:"cast"`
}

// Save inserts a list of movies into a storage that will be used by the search API
func (storage *Storage) Save(movies []Movie) error {
	moviesDTO := make([]movieDTO, 0, len(movies))

	for _, movie := range movies {
		movieDTO := movieDTO{
			ID:       movie.ID,
			Name:     movie.Name,
			Director: movie.Director,
			Cast:     movie.Cast,
		}

		moviesDTO = append(moviesDTO, movieDTO)
	}

	createMoviesRequest := createMoviesRequest{
		Movies: moviesDTO,
	}

	return storage.createMovies(createMoviesRequest)
}

func (storage *Storage) createMovies(createMoviesRequest createMoviesRequest) error {
	request, err := storage.getCreateMoviesHTTPRequest(createMoviesRequest)
	if err != nil {
		return err
	}

	return retry.Do(func() error {
		return executeRequest(request)
	}, retry.Attempts(3))
}

func (storage *Storage) getCreateMoviesHTTPRequest(createMoviesRequest createMoviesRequest) (*http.Request, error) {
	endpointURL := fmt.Sprintf("%s/api/movies", storage.options.StorageURL)

	requestBytes, err := json.Marshal(createMoviesRequest)
	if err != nil {
		return nil, fmt.Errorf("Could not send request to storage: %w", err)
	}

	request, err := http.NewRequest("POST", endpointURL, bytes.NewBuffer(requestBytes))
	request.Header.Set("Content-Type", "application/json")

	return request, nil
}

func executeRequest(request *http.Request) error {
	client := &http.Client{}
	response, err := client.Do(request)

	if err != nil {
		return fmt.Errorf("Could not execute request: %w", err)
	}

	defer response.Body.Close()
	if response.StatusCode != 201 {
		return fmt.Errorf("Response obtained from storage is unexpected. Expected 201, got %d", response.StatusCode)
	}

	return nil
}
