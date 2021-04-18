package movie_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/mathnogueira/imdb-api/storage/movie"
)

var _ = Describe("DatabaseItem", func() {

	It("Should generate all keys for a movie", func() {
		jurassicPark := movie.Movie{
			Name:     "Jurassic Park",
			Director: "Steven Spielberg",
			Cast:     []string{"Sam Neil", "Jeff Goldblum"},
		}

		databaseItem := movie.DatabaseItem{jurassicPark}

		keys := databaseItem.GetKeys()

		Expect(len(keys)).To(Equal(8))
		Expect(keys).To(ContainElement("Jurassic"))
		Expect(keys).To(ContainElement("Park"))
		Expect(keys).To(ContainElement("Steven"))
		Expect(keys).To(ContainElement("Spielberg"))
		Expect(keys).To(ContainElement("Sam"))
		Expect(keys).To(ContainElement("Neil"))
		Expect(keys).To(ContainElement("Jeff"))
		Expect(keys).To(ContainElement("Goldblum"))
	})
})
