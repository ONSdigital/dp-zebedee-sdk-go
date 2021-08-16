package zebedee

import (
	"context"
	"fmt"
	dphttp "github.com/ONSdigital/dp-net/http"
	"io/ioutil"
	"net/http"
	"time"
)

const (
	incorrectStatusErrFmt         = "request %s %s expected status %d but received %d"
	incorrectStatusWithBodyErrFmt = "request %s %s expected status %d but received %d response body: %s"
)

//go:generate moq -out mock/httpclient.go -pkg mock . HttpClient
// HttpClient defines a Zebedee HTTP client
type HttpClient interface {
	Do(ctx context.Context, req *http.Request) (*http.Response, error)
}

//NewHttpClient Construct a new HttpClient
func NewHttpClient(timeout time.Duration) HttpClient {
	return dphttp.ClientWithTimeout(nil, timeout)
}

//checkResponseStatus return an error if the actual response status did not match the expected.
func checkResponseStatus(resp *http.Response, expected int) error {
	req := resp.Request
	if resp.StatusCode != expected {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return fmt.Errorf("unexpected error attempting to read error response body: %s", err.Error())
		}

		if len(body) > 0 {
			return fmt.Errorf(incorrectStatusWithBodyErrFmt, req.Method, req.URL.RequestURI(), expected, resp.StatusCode, string(body))
		}

		return fmt.Errorf(incorrectStatusErrFmt, req.Method, req.URL.RequestURI(), expected, resp.StatusCode)
	}
	return nil
}
