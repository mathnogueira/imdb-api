package movie_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"go.uber.org/zap"

	"github.com/mathnogueira/imdb-api/storage/database"
	"github.com/mathnogueira/imdb-api/storage/movie"
)

var _ = Describe("Persistence", func() {

	var repository *movie.Repository
	logger := zap.NewNop()

	BeforeEach(func() {
		storage := database.NewMemoryStorage(logger)
		repository = movie.NewRepository(storage, logger)

		movies := []movie.Movie{
			{Name: "Jurassic Park", Director: "Steven Spielberg", Cast: []string{"Sam Neil", "Jeff Goldblum"}},
			{Name: "Jaws", Director: "Steven Spielberg", Cast: []string{"Roy Schneider", "Robert Shaw"}},
			{Name: "Monty Pithon's Life of Brian", Director: "Terry Jones", Cast: []string{"Graham Chapman", "John Cleese"}},
		}

		for _, movie := range movies {
			repository.Save(movie)
		}
	})

	It("Should save and retrieve movies from storage", func() {
		spielbergMovies := repository.Search([]string{"spielberg"})

		Expect(spielbergMovies).To(HaveLen(2))
		movieNames := []string{
			spielbergMovies[0].Name,
			spielbergMovies[1].Name,
		}

		Expect(movieNames).Should(ContainElement("Jurassic Park"))
		Expect(movieNames).Should(ContainElement("Jaws"))
	})
})
