package transformation_test

import (
	"github.com/mathnogueira/imdb-api/webcrawler/imdb"
	"github.com/mathnogueira/imdb-api/webcrawler/transformation"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("ImdbModelToMovie", func() {

	var theGodfatherMovie imdb.Movie
	var malformedMovie imdb.Movie

	BeforeEach(func() {
		theGodfatherMovie = imdb.Movie{
			Name: "The Godfather",
			URL:  "/title/tt0068646/?ref_=adv_li_tt",
			Cast: "Francis Ford Coppola (dir.), Marlon Brando, Al Pacino",
		}

		malformedMovie = imdb.Movie{
			Name: "Malformed movie",
			URL:  "title/tt0068632/?ref_=adv_li_tt",
			Cast: "Nobody famous, That guy from Jurassic Park, Benedict Cucumber",
		}
	})

	It("Should keep the same name", func() {
		movie, err := transformation.ConvertIMDBMovieToMovie(theGodfatherMovie)

		Expect(err).ShouldNot(HaveOccurred())
		Expect(movie.Name).To(Equal("The Godfather"))
	})

	It("Should extract the movie ID from IMDB URL", func() {
		movie, err := transformation.ConvertIMDBMovieToMovie(theGodfatherMovie)

		Expect(err).ShouldNot(HaveOccurred())
		Expect(movie.ID).To(Equal("0068646"))
	})

	It("Should return an error when the URL is malformed and an id cannot be extracted", func() {
		_, err := transformation.ConvertIMDBMovieToMovie(malformedMovie)

		Expect(err).Should(HaveOccurred())
	})

	It("Should extract director from cast string", func() {
		movie, err := transformation.ConvertIMDBMovieToMovie(theGodfatherMovie)

		Expect(err).ShouldNot(HaveOccurred())
		Expect(movie.Director).To(Equal("Francis Ford Coppola"))
	})

	It("Should extract the cast from the movie", func() {
		movie, err := transformation.ConvertIMDBMovieToMovie(theGodfatherMovie)

		Expect(err).ShouldNot(HaveOccurred())
		Expect(len(movie.Cast)).To(Equal(2))
		Expect(movie.Cast[0]).To(Equal("Marlon Brando"))
		Expect(movie.Cast[1]).To(Equal("Al Pacino"))
	})
})
