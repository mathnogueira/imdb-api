package extractor

import (
	"fmt"

	"github.com/mathnogueira/imdb-api/webcrawler/imdb"
	"github.com/mathnogueira/imdb-api/webcrawler/movie"
	"github.com/mathnogueira/imdb-api/webcrawler/transformation"
	"go.uber.org/zap"
)

// Extractor is responsible for getting information from IMDB and sending it
// to the storage API
type Extractor struct {
	logger *zap.Logger
}

// NewExtractor creates a new instance of a data extractor
func NewExtractor(logger *zap.Logger) *Extractor {
	return &Extractor{logger}
}

// Execute the crawler routine
func (extractor *Extractor) Execute(options Options) error {
	imdbCrawler := imdb.NewCrawler(extractor.logger)
	storageOptions := movie.StorageOptions{
		StorageURL: options.StorageURL,
	}
	storage := movie.NewStorage(storageOptions)

	extractor.logger.Debug("Executing extraction", zap.String("storageURL", options.StorageURL))

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
