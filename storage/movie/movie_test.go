package movie_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/mathnogueira/imdb-api/storage/movie"
)

var _ = Describe("DatabaseItem", func() {

	It("Should generate all keys for a movie as lowercase", func() {
		jurassicPark := movie.Movie{
			Name:     "Jurassic Park",
			Director: "Steven Spielberg",
			Cast:     []string{"Sam Neil", "Jeff Goldblum"},
		}

		databaseItem := movie.DatabaseItem{jurassicPark}

		keys := databaseItem.GetKeys()

		Expect(len(keys)).To(Equal(8))
		Expect(keys).To(ContainElement("jurassic"))
		Expect(keys).To(ContainElement("park"))
		Expect(keys).To(ContainElement("steven"))
		Expect(keys).To(ContainElement("spielberg"))
		Expect(keys).To(ContainElement("sam"))
		Expect(keys).To(ContainElement("neil"))
		Expect(keys).To(ContainElement("jeff"))
		Expect(keys).To(ContainElement("goldblum"))
	})
})
