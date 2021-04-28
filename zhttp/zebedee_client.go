package zhttp

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

// Client defines a Zebedee HTTP client
type Client interface {
	Do(r *http.Request) (*http.Response, error)
	RequestObject(r *http.Request, expectedStatus int, entity interface{}) error
	GetHost() string
	NewAuthenticatedRequest(uri, authToken, method string, entity interface{}) (*http.Request, error)
}

type zebedeeClient struct {
	Host    string
	httpCli *http.Client
}

func NewClient(host string) Client {
	return &zebedeeClient{
		Host:    host,
		httpCli: &http.Client{Timeout: time.Second * 3},
	}
}

func (c *zebedeeClient) RequestObject(r *http.Request, expectedStatus int, entity interface{}) error {
	resp, err := c.Do(r)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != expectedStatus {
		return IncorrectStatusErr(r.RequestURI, r.Method, expectedStatus, resp.StatusCode)
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(b, &entity); err != nil {
		return err
	}

	return nil
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

	url := fmt.Sprintf("%s%s", c.Host, uri)
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}

	req.Header.Set("content-type", "application/json")
	req.Header.Set("X-Florence-Token", authToken)
	return req, nil
}

func (c *zebedeeClient) Do(r *http.Request) (*http.Response, error) {
	return c.httpCli.Do(r)
}

func (c *zebedeeClient) GetHost() string {
	return c.Host
}

func IncorrectStatusErr(endpoint, method string, expected, actual int) error {
	return fmt.Errorf("%s %s returned status %d but expected %d", expected, method, expected, actual)
}
