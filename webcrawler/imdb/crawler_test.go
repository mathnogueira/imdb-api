package imdb_test

import (
	"github.com/mathnogueira/imdb-api/webcrawler/imdb"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"go.uber.org/zap"
)

var _ = Describe("IMDB Crawler", func() {

	logger := zap.NewNop()

	It("Should be able to parse movies from IMDB page", func() {
		crawler := imdb.NewCrawler(logger)
		extractedMovies := crawler.GetTopMovies()

		Expect(len(extractedMovies)).To(Equal(1000))

		for _, movie := range extractedMovies {
			Expect(len(movie.Name)).ToNot(Equal(0))
			Expect(len(movie.URL)).ToNot(Equal(0))
			Expect(len(movie.Cast)).ToNot(Equal(0))
		}
	})
})
