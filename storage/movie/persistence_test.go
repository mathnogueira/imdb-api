package movie_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/mathnogueira/imdb-api/storage/database"
	"github.com/mathnogueira/imdb-api/storage/movie"
)

var _ = Describe("Persistence", func() {

	var repository *movie.Repository

	BeforeEach(func() {
		storage := database.NewMemoryStorage()
		repository = movie.NewRepository(storage)

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
		Expect(spielbergMovies[0].Name).To(Equal("Jurassic Park"))
		Expect(spielbergMovies[1].Name).To(Equal("Jaws"))
	})
})
