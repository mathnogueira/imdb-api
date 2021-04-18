package testing

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

// Request represents a HTTP request that will be sent to the server during tests
type Request struct {
	Method  string
	URL     string
	Payload interface{}
}

// Response is what the HTTP returns after a request
type Response struct {
	StatusCode int
}

// SendRequest sends a request to a server and populates the response
func SendRequest(request Request, response interface{}) (*Response, error) {
	httpRequest, err := createHTTPRequest(request)
	if err != nil {
		return nil, err
	}

	client := http.Client{}
	httpResponse, err := client.Do(httpRequest)

	if err != nil {
		return nil, fmt.Errorf("Could not execute request: %w", err)
	}

	responseBodyBytes, err := ioutil.ReadAll(httpResponse.Body)
	if err != nil {
		return nil, fmt.Errorf("Could not read response body: %w", err)
	}

	if len(responseBodyBytes) > 0 {
		err = json.Unmarshal(responseBodyBytes, response)
		if err != nil {
			return nil, fmt.Errorf("Could not parse response body into JSON: %w", err)
		}
	}

	return &Response{StatusCode: httpResponse.StatusCode}, nil
}

func createHTTPRequest(request Request) (*http.Request, error) {
	requestURL, err := url.Parse(request.URL)
	if err != nil {
		return nil, fmt.Errorf("Could not parse URL: %w", err)
	}

	payloadBytes, err := json.Marshal(request.Payload)
	if err != nil {
		return nil, fmt.Errorf("Could not parse JSON content: %w", err)
	}

	jsonReader := ioutil.NopCloser(bytes.NewReader(payloadBytes))

	httpRequest := &http.Request{
		Method: request.Method,
		URL:    requestURL,
		Body:   jsonReader,
	}

	return httpRequest, nil
}
