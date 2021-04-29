package zebedee

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

const (
	incorrectStatusErrFmt = "%s %s expected status %d but received %d"
)

// HttpClient defines a Zebedee HTTP client
type HttpClient interface {
	Do(r *http.Request) (*http.Response, error)
	RequestObject(r *http.Request, expectedStatus int, entity interface{}) error
}

type httpClient struct {
	httpCli *http.Client
}

//NewHttpClient Construct a new HttpClient
func NewHttpClient(timeout time.Duration) HttpClient {
	return &httpClient{
		httpCli: &http.Client{Timeout: timeout},
	}
}

//RequestObject execute a JSON http request and unmarshal the response into the provided entity
func (c *httpClient) RequestObject(r *http.Request, expectedStatus int, entity interface{}) error {
	resp, err := c.Do(r)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if err = checkResponseStatus(resp, expectedStatus); err != nil {
		return err
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

//Do execute the http request
func (c *httpClient) Do(r *http.Request) (*http.Response, error) {
	return c.httpCli.Do(r)
}

//checkResponseStatus return an error if the actual response status did not match the expected.
func checkResponseStatus(resp *http.Response, expected int) error {
	req := resp.Request
	if resp.StatusCode != expected {
		return fmt.Errorf(incorrectStatusErrFmt, req.RequestURI, req.Method, expected, resp.StatusCode)
	}
	return nil
}

//discardResponse consume the response body and send it to dev/null
func discardResponse(resp *http.Response) error {
	_, err := io.Copy(ioutil.Discard, resp.Body)
	return err
}
