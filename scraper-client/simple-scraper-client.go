package client

import (
	"log"
	"net/http"
	"net/url"
	"os"
	"time"

	generator "github.com/alexmorten/smag-mvp/http_header-generator"
)

// SimpleScraperClient handles retries and setting random headers for scraping
type SimpleScraperClient struct {
	currentAddress string
	client         *http.Client
	instanceID     string
	*generator.HTTPHeaderGenerator
}

// NewSimpleScraperClient returns an initialized SimpleScraperClient
func NewSimpleScraperClient() *SimpleScraperClient {
	client := &SimpleScraperClient{}
	client.HTTPHeaderGenerator = generator.New()

	path := os.Getenv("HTTP_PROXY")
	if path != "" {
		proxy, err := url.Parse(path)
		if err != nil {
			log.Fatalf("failed to parse proxy URL '%s' with '%v'", path, err)
		}
		client.client = &http.Client{
			Transport: &http.Transport{Proxy: http.ProxyURL(proxy)},
			Timeout:   time.Second * 18,
		}
	} else {
		client.client = &http.Client{}
	}

	return client
}

// WithRetries calls f with retries
func (s *SimpleScraperClient) WithRetries(times int, f func() error) error {
	var err error
	for i := 0; i < times; i++ {
		err = f()

		if err == nil {
			return nil
		}
		time.Sleep(100 * time.Millisecond)
	}
	return err
}

// Do the request with correct headers
func (s *SimpleScraperClient) Do(request *http.Request) (*http.Response, error) {
	s.AddHeaders(&request.Header)
	return s.client.Do(request)
}
