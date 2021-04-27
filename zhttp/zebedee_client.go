package zhttp

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

// Client defines a Zebedee HTTP client
type Client interface {
	Do(r *http.Request) (*http.Response, error)
	NewAuthenticatedRequest(uri, authToken, method string, entity interface{}) (*http.Request, error)
}

type zebedeeClient struct {
	host    string
	httpCli *http.Client
}

func NewClient(host string) Client {
	return &zebedeeClient{
		host:    host,
		httpCli: &http.Client{Timeout: time.Second * 3},
	}
}

func (c *zebedeeClient) NewAuthenticatedRequest(uri, authToken, method string, entity interface{}) (*http.Request, error) {
	var body io.Reader
	if entity != nil {
		b, err := json.Marshal(entity)
		if err != nil {
			return nil, err
		}

		body = bytes.NewReader(b)
	}

	url := fmt.Sprintf("%s%s", c.host, uri)
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}

	req.Header.Set("content-type", "application/json")
	req.Header.Set("X-Florence-Token", authToken)
	return req, nil
}

func (c *zebedeeClient) Do(r *http.Request) (*http.Response, error) {
	return c.Do(r)
}
