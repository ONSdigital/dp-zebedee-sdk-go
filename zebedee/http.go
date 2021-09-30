package zebedee

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	dphttp "github.com/ONSdigital/dp-net/http"
)

const (
	readResponseBodyErrFmt        = "unexpected error attempting to read error response body: %q %s %s expected status: %d actual status: %d"
	incorrectStatusErrFmt         = "request %s %s expected status %d but received %d"
	incorrectStatusWithBodyErrFmt = "request %s %s expected status %d but received %d"
)

//go:generate moq -out mock/httpclient.go -pkg mock . HttpClient
// HttpClient defines a Zebedee HTTP client
type HttpClient interface {
	Do(ctx context.Context, req *http.Request) (*http.Response, error)
}

// APIError represent an error returned from the Zebedee CMS API.
type APIError struct {
	ActualStatus   int
	ExpectedStatus int
	Message        string
	Body           string
}

func (err *APIError) Error() string {
	return err.Message
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
			return &APIError{
				ActualStatus:   resp.StatusCode,
				ExpectedStatus: expected,
				Message:        fmt.Sprintf(readResponseBodyErrFmt, err.Error(), req.Method, req.URL.RequestURI(), expected, resp.StatusCode),
			}
		}

		if len(body) > 0 {
			return &APIError{
				ActualStatus:   resp.StatusCode,
				ExpectedStatus: expected,
				Body:           string(body),
				Message:        fmt.Sprintf(incorrectStatusWithBodyErrFmt, req.Method, req.URL.RequestURI(), expected, resp.StatusCode),
			}
		}

		return &APIError{
			ActualStatus:   resp.StatusCode,
			ExpectedStatus: expected,
			Message:        fmt.Sprintf(incorrectStatusErrFmt, req.Method, req.URL.RequestURI(), expected, resp.StatusCode),
		}
	}
	return nil
}
