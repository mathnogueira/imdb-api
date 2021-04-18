package movie_test

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"sync"
)

type createMoviesRequest struct {
	Movies []struct {
		ID       string   `json:"id"`
		Name     string   `json:"name"`
		Director string   `json:"director"`
		Cast     []string `json:"cast"`
	} `json:"movies"`
}

func setupStorageMockServer() *httptest.Server {
	var server *httptest.Server
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		server = httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, req *http.Request) {
			requestBody, err := ioutil.ReadAll(req.Body)
			if err != nil {
				writer.WriteHeader(500)
				return
			}

			var storageRequest createMoviesRequest
			err = json.Unmarshal(requestBody, &storageRequest)

			if err != nil {
				writer.WriteHeader(500)
				return
			}

			for _, movie := range storageRequest.Movies {
				if movie.ID == "" || movie.Director == "" || movie.Name == "" || len(movie.Cast) == 0 {
					writer.WriteHeader(400)
					return
				}

				for _, actor := range movie.Cast {
					if actor == "" {
						writer.WriteHeader(400)
						return
					}
				}
			}

			writer.WriteHeader(201)
		}))
		wg.Done()
	}()

	wg.Wait()

	return server
}
