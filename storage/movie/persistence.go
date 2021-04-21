package movie

import (
	"fmt"

	"github.com/mathnogueira/imdb-api/storage/database"
	"go.uber.org/zap"
)

// Repository abstracts the movie storage from the rest of the application
type Repository struct {
	storage database.Storage
	logger  *zap.Logger
}

// NewRepository creates a new repository for handling persistence operations of movies
func NewRepository(storage database.Storage, logger *zap.Logger) *Repository {
	return &Repository{storage: storage, logger: logger}
}

// Save a movie in the storage
func (repository *Repository) Save(movie Movie) {
	repository.logger.Debug("Saving new movie",
		zap.String("id", movie.ID),
		zap.String("name", movie.Name),
	)
	databaseItem := DatabaseItem{Content: movie}

	repository.storage.Save(databaseItem)
}

// Search all movies and return the ones that match all the provided keys
func (repository *Repository) Search(keys []string) []Movie {
	repository.logger.Debug("Searching movies", zap.Strings("keys", keys))
	databaseItems := repository.storage.Search(keys)

	movies := make([]Movie, 0, len(databaseItems))

	for _, databaseItem := range databaseItems {
		movie := databaseItem.GetContent().(Movie)
		movies = append(movies, movie)
	}

	repository.logger.Debug(fmt.Sprintf("%d movies found", len(movies)))

	return movies
}
