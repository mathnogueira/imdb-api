package movie

import "github.com/mathnogueira/imdb-api/storage/database"

// Repository abstracts the movie storage from the rest of the application
type Repository struct {
	storage database.Storage
}

// NewRepository creates a new repository for handling persistence operations of movies
func NewRepository(storage database.Storage) *Repository {
	return &Repository{storage: storage}
}

// Save a movie in the storage
func (repository *Repository) Save(movie Movie) {
	databaseItem := DatabaseItem{Content: movie}

	repository.storage.Save(databaseItem)
}

// Search all movies and return the ones that match all the provided keys
func (repository *Repository) Search(keys []string) []Movie {
	databaseItems := repository.storage.Search(keys)

	movies := make([]Movie, 0, len(databaseItems))

	for _, databaseItem := range databaseItems {
		movie := databaseItem.GetContent().(Movie)
		movies = append(movies, movie)
	}

	return movies
}
