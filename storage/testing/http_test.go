package testing_test

import (
	"net/http"
	"net/http/httptest"
	"sync"

	"github.com/mathnogueira/imdb-api/storage/testing"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

type request struct {
	FirstNumber  int `json:"first_number"`
	SecondNumber int `json:"second_number"`
}

type response struct {
	Sum int `json:"sum"`
}

var _ = Describe("Http", func() {

	var server *httptest.Server
	var emptyResponseServer *httptest.Server

	BeforeSuite(func() {
		var wg sync.WaitGroup
		wg.Add(1)
		go func() {
			server = httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, req *http.Request) {
				responseJSON := `
				{
					"sum": 18
				}
				`
				writer.WriteHeader(http.StatusCreated)
				writer.Write([]byte(responseJSON))
			}))

			emptyResponseServer = httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, req *http.Request) {
				writer.WriteHeader(http.StatusOK)
			}))

			wg.Done()
		}()

		wg.Wait()
	})

	AfterSuite(func() {
		server.Close()
		emptyResponseServer.Close()
	})

	It("Should be able to send a request to a server and retrieve its response", func() {
		payload := request{
			FirstNumber:  8,
			SecondNumber: 10,
		}

		var response response

		request := testing.Request{
			Method:  "GET",
			URL:     server.URL,
			Payload: payload,
		}
		httpResponse, err := testing.SendRequest(request, &response)

		Expect(err).ShouldNot(HaveOccurred())
		Expect(response.Sum).To(Equal(18))
		Expect(httpResponse.StatusCode).To(Equal(http.StatusCreated))
	})

	It("Should not parse response body if it's empty", func() {
		payload := request{
			FirstNumber:  8,
			SecondNumber: 10,
		}

		var response response

		request := testing.Request{
			Method:  "GET",
			URL:     emptyResponseServer.URL,
			Payload: payload,
		}
		httpResponse, err := testing.SendRequest(request, &response)

		Expect(err).ShouldNot(HaveOccurred())
		Expect(response.Sum).To(Equal(0))
		Expect(httpResponse.StatusCode).To(Equal(http.StatusOK))
	})
})
