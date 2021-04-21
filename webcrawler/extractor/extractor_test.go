package extractor_test

import (
	"net/http"
	"net/http/httptest"
	"sync"

	"github.com/mathnogueira/imdb-api/webcrawler/extractor"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"go.uber.org/zap"
)

var _ = Describe("Extractor", func() {

	logger := zap.NewNop()

	It("Should execute without any error when storage API is online", func() {
		server := setupMockServer()
		defer server.Close()

		imdbExtractor := extractor.NewExtractor(logger)
		extractorOptions := extractor.Options{StorageURL: server.URL}
		err := imdbExtractor.Execute(extractorOptions)

		Expect(err).ShouldNot(HaveOccurred())
	})

	It("Should return an error if storage API is offline", func() {
		server := setupMockServer()

		imdbExtractor := extractor.NewExtractor(logger)
		extractorOptions := extractor.Options{StorageURL: server.URL}
		// Forces the server to close before the execution of the crawler
		// This simulates downtime in the storage API
		server.Close()

		err := imdbExtractor.Execute(extractorOptions)

		Expect(err).Should(HaveOccurred())
	})
})

func setupMockServer() *httptest.Server {
	var server *httptest.Server
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		server = httptest.NewServer(http.HandlerFunc(func(writer http.ResponseWriter, req *http.Request) {
			writer.WriteHeader(201)
		}))
		wg.Done()
	}()

	wg.Wait()
	return server
}
