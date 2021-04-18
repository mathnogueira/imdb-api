package transformation

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/mathnogueira/imdb-api/webcrawler/imdb"
	"github.com/mathnogueira/imdb-api/webcrawler/movie"
)

var movieIDRegex = regexp.MustCompile("\\/title/tt([0-9]+)/*")

// ConvertIMDBMovieToMovie receives a movie extracted from IMDB and transforms it
// into a domain object used across the whole project.
func ConvertIMDBMovieToMovie(imdbMovie imdb.Movie) (movie.Movie, error) {
	movieIDResult := movieIDRegex.FindStringSubmatch(imdbMovie.URL)
	if len(movieIDResult) < 2 {
		return movie.Movie{}, fmt.Errorf("Movie %s has a malformed URL and cannot be processed", imdbMovie.Name)
	}

	director, actors := extractCastFromMovie(imdbMovie)

	return movie.Movie{
		Name:     imdbMovie.Name,
		ID:       movieIDResult[1],
		Director: director,
		Cast:     actors,
	}, nil
}

// extractCastFromMovie receives an IMDB Movie and extracts its director and actors. The
// director is returned as the first return arg from the function as the actor are returned
// in the second argument.
func extractCastFromMovie(imdbMovie imdb.Movie) (string, []string) {
	cast := strings.Split(imdbMovie.Cast, ", ")

	// Director is always the first one
	directorName := cast[0]
	// This removes the suffix " (dir.) from the director name"
	directorName = directorName[:len(directorName)-7]

	return directorName, cast[1:]
}
