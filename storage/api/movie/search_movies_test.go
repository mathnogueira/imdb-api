package movie_test

import (
	"fmt"
	"net/http"

	"github.com/mathnogueira/imdb-api/storage/api"
	"github.com/mathnogueira/imdb-api/storage/testing"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

type searchMoviesRequest struct {
	Terms []string `json:"terms"`
}

type searchMoviesResponse struct {
	Movies []movieDTO
}

var _ = Describe("SearchMovies", func() {

	var server *api.Server

	BeforeEach(func() {
		server = api.NewServer(8888)
		go server.Start()
		populateDatabaseWithMovies(server)
	})

	AfterEach(func() {
		server.Close()
	})

	It("Should return movies from Steven Spielberg", func() {
		response, httpResponse, err := searchMovies(server, []string{"spielberg"})

		Expect(err).ShouldNot(HaveOccurred())
		Expect(httpResponse.StatusCode).To(Equal(http.StatusOK))
		Expect(response.Movies).To(HaveLen(2))
	})

	It("Should return an empty array when no movie matches the terms", func() {
		response, httpResponse, err := searchMovies(server, []string{"spielberg", "Godfather"})

		Expect(err).ShouldNot(HaveOccurred())
		Expect(httpResponse.StatusCode).To(Equal(http.StatusOK))
		Expect(response.Movies).To(HaveLen(0))
	})

	It("Should return Jaws when we search for spielberg and Schneider", func() {
		response, httpResponse, err := searchMovies(server, []string{"spielberg", "Schneider"})

		Expect(err).ShouldNot(HaveOccurred())
		Expect(httpResponse.StatusCode).To(Equal(http.StatusOK))
		Expect(response.Movies).To(HaveLen(1))
	})
})

func populateDatabaseWithMovies(server *api.Server) {
	movies := []movieDTO{
		{
			ID:       "00123456",
			Name:     "Jurassic Park",
			Director: "Steven Spielberg",
			Cast:     []string{"Sam Neil", "Jeff Goldblum"},
		},
		{
			ID:       "00654321",
			Name:     "Jaws",
			Director: "Steven Spielberg",
			Cast:     []string{"Roy Schneider", "Robert Shaw"},
		},
		{
			ID:       "0068646",
			Name:     "The Godfather",
			Director: "Francis Ford Coppola",
			Cast:     []string{"Marlon Brando", "Al Pacino"},
		},
	}

	_, err := insertMovies(movies, server)
	Expect(err).ShouldNot(HaveOccurred())
}

func searchMovies(server *api.Server, terms []string) (searchMoviesResponse, *testing.Response, error) {
	requestPayload := searchMoviesRequest{
		Terms: terms,
	}
	request := testing.Request{
		Method:  "POST",
		URL:     fmt.Sprintf("%s/api/movies/search", server.GetAddress()),
		Payload: requestPayload,
	}

	var response searchMoviesResponse
	httpResponse, err := testing.SendRequest(request, &response)

	return response, httpResponse, err
}
