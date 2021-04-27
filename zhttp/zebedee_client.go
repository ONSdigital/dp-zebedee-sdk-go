package zhttp

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

// Client defines a Zebedee HTTP client
type Client interface {
	Do(r *http.Request) (*http.Response, error)
}

func NewAuthenticatedRequest(url, authToken, method string, entity interface{}) (*http.Request, error) {
	var body io.Reader
	if entity != nil {
		b, err := json.Marshal(entity)
		if err != nil {
			return nil, err
		}

		body = bytes.NewReader(b)
	}

	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}

	req.Header.Set("content-type", "application/json")
	req.Header.Set("X-Florence-Token", authToken)
	return req, nil
}
