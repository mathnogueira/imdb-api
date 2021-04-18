package movie_test

import (
	"github.com/mathnogueira/imdb-api/webcrawler/movie"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Storage", func() {

	It("Should send a POST request when saving the movies into the storage", func() {
		server := setupStorageMockServer()
		defer server.Close()

		storageOptions := movie.StorageOptions{
			StorageURL: server.URL,
		}
		storage := movie.NewStorage(storageOptions)

		movies := []movie.Movie{
			{
				ID:       "0079470",
				Name:     "Monty Python's Life of Brian",
				Director: "Terry Jones",
				Cast:     []string{"Graham Chapman", "John Cleese"},
			},
			{
				ID:       "0118715",
				Name:     "The Big Lebowski",
				Director: "Joel Coen",
				Cast:     []string{"Jeff Bridges", "John Goodman"},
			},
		}

		err := storage.Save(movies)
		Expect(err).ShouldNot(HaveOccurred())
	})
})
