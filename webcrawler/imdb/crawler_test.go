package imdb_test

import (
	"fmt"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/mathnogueira/imdb-api/webcrawler/imdb"
)

var _ = Describe("IMDB Crawler", func() {

	It("Should be able to parse movies from IMDB page", func() {
		options := imdb.CrawlerOptions{}

		extractedMovies := extractMovies(options)

		Expect(len(extractedMovies)).To(Equal(1000))

		for _, movie := range extractedMovies {
			Expect(len(movie.Name)).ToNot(Equal(0))
			Expect(len(movie.URL)).ToNot(Equal(0))
			Expect(len(movie.Cast)).ToNot(Equal(0))
		}
	})
})

func extractMovies(options imdb.CrawlerOptions) []imdb.Movie {
	movieChannel := make(chan imdb.Movie)
	doneChannel := make(chan bool)
	movies := make([]imdb.Movie, 0)

	crawler := imdb.NewCrawler(options)
	go crawler.Start(movieChannel, doneChannel)

loop:
	for {
		select {
		case movie := <-movieChannel:
			movies = append(movies, movie)
		case <-doneChannel:
			break loop
		}
	}

	for _, movie := range movies {
		fmt.Println(movie)
	}

	return movies
}
