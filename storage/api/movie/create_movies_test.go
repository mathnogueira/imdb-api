package movie_test

import (
	"fmt"
	"net/http"

	"github.com/mathnogueira/imdb-api/storage/api"
	"github.com/mathnogueira/imdb-api/storage/testing"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

type createMoviesRequest struct {
	Movies []movieDTO `json:"movies"`
}

type movieDTO struct {
	ID       string   `json:"id"`
	Name     string   `json:"name"`
	Director string   `json:"director"`
	Cast     []string `json:"cast"`
}

type response struct{}

var _ = Describe("CreateMovies", func() {

	var server *api.Server

	BeforeSuite(func() {
		server = api.NewServer(9876)
		go server.Start()
	})

	AfterSuite(func() {
		err := server.Close()
		Expect(err).ShouldNot(HaveOccurred())
	})

	It("Should return status 201 when a valid JSON is provided", func() {
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
				Cast:     []string{"Roy Scheider", "Robert Shaw"},
			},
		}

		httpResponse, err := insertMovies(movies, server)

		Expect(err).ShouldNot(HaveOccurred())
		Expect(httpResponse.StatusCode).To(Equal(http.StatusCreated))
	})

	It("Should return status 400 if movie ID is missing", func() {
		movies := []movieDTO{
			{
				Name:     "Jurassic Park",
				Director: "Steven Spielberg",
				Cast:     []string{"Sam Neil", "Jeff Goldblum"},
			},
		}

		httpResponse, err := insertMovies(movies, server)

		Expect(err).ShouldNot(HaveOccurred())
		Expect(httpResponse.StatusCode).To(Equal(http.StatusBadRequest))
	})

	It("Should return status 400 if movie Name is missing", func() {
		movies := []movieDTO{
			{
				ID:       "00123456",
				Director: "Steven Spielberg",
				Cast:     []string{"Sam Neil", "Jeff Goldblum"},
			},
		}

		httpResponse, err := insertMovies(movies, server)

		Expect(err).ShouldNot(HaveOccurred())
		Expect(httpResponse.StatusCode).To(Equal(http.StatusBadRequest))
	})

	It("Should return status 400 if movie Director is missing", func() {
		movies := []movieDTO{
			{
				ID:   "00123456",
				Name: "Jurassic Park",
				Cast: []string{"Sam Neil", "Jeff Goldblum"},
			},
		}

		httpResponse, err := insertMovies(movies, server)

		Expect(err).ShouldNot(HaveOccurred())
		Expect(httpResponse.StatusCode).To(Equal(http.StatusBadRequest))
	})

	It("Should return status 400 if movie Cast member name is empty", func() {
		movies := []movieDTO{
			{
				ID:       "00123456",
				Name:     "Jurassic Park",
				Director: "Steven Spielberg",
				Cast:     []string{""},
			},
		}

		httpResponse, err := insertMovies(movies, server)

		Expect(err).ShouldNot(HaveOccurred())
		Expect(httpResponse.StatusCode).To(Equal(http.StatusBadRequest))
	})
})

func insertMovies(movies []movieDTO, server *api.Server) (*testing.Response, error) {
	moviesJSON := createMoviesRequest{
		Movies: movies,
	}

	request := testing.Request{
		Method:  "POST",
		URL:     fmt.Sprintf("%s/api/movies", server.GetAddress()),
		Payload: moviesJSON,
	}

	var response response
	httpResponse, err := testing.SendRequest(request, &response)

	return httpResponse, err
}
