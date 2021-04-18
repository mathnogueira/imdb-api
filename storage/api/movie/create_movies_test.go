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
		go func() {
			server.Start()
		}()
	})

	AfterSuite(func() {
		err := server.Close()
		Expect(err).ShouldNot(HaveOccurred())
	})

	It("Should return status 201 when a valid JSON is provided", func() {
		moviesJSON := createMoviesRequest{
			Movies: []movieDTO{
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
			},
		}

		request := testing.Request{
			Method:  "POST",
			URL:     fmt.Sprintf("%s/api/movies", server.GetAddress()),
			Payload: moviesJSON,
		}

		var response response
		httpResponse, err := testing.SendRequest(request, &response)

		Expect(err).ShouldNot(HaveOccurred())
		Expect(httpResponse.StatusCode).To(Equal(http.StatusCreated))
	})
})
