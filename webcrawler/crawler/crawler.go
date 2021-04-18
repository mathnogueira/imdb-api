package crawler

import (
	"fmt"

	"github.com/mathnogueira/imdb-api/webcrawler/imdb"
	"github.com/mathnogueira/imdb-api/webcrawler/movie"
	"github.com/mathnogueira/imdb-api/webcrawler/transformation"
)

// Execute the crawler routine
func Execute(options Options) error {
	imdbCrawler := imdb.NewCrawler()
	storageOptions := movie.StorageOptions{
		StorageURL: options.StorageURL,
	}
	storage := movie.NewStorage(storageOptions)

	imdbMovies := imdbCrawler.GetTopMovies()
	movies := make([]movie.Movie, 0, len(imdbMovies))

	for _, imdbMovie := range imdbMovies {
		movie, err := transformation.ConvertIMDBMovieToMovie(imdbMovie)
		if err != nil {
			return fmt.Errorf("Could not convert IMDB Movie to a Movie: %w", err)
		}
		movies = append(movies, movie)
	}

	return storage.Save(movies)
}
