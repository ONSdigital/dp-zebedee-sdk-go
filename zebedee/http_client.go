package zebedee

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// HttpClient defines a Zebedee HTTP client
type HttpClient interface {
	Do(r *http.Request) (*http.Response, error)
	RequestObject(r *http.Request, expectedStatus int, entity interface{}) error
}

type httpClient struct {
	//Host    string
	httpCli *http.Client
}

func NewHttpClient() HttpClient {
	return &httpClient{
		httpCli: &http.Client{Timeout: time.Second * 3},
	}
}

func (c *httpClient) RequestObject(r *http.Request, expectedStatus int, entity interface{}) error {
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

func (c *httpClient) Do(r *http.Request) (*http.Response, error) {
	return c.httpCli.Do(r)
}

func IncorrectStatusErr(endpoint, method string, expected, actual int) error {
	return fmt.Errorf("%s %s expected status %d but received %d", endpoint, method, expected, actual)
}
