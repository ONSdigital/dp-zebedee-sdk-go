package zebedee

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/ONSdigital/dp-net/v2/request"
	"github.com/ONSdigital/dp-zebedee-sdk-go/zebedee/mock"
	. "github.com/smartystreets/goconvey/convey"
)

var (
	testSess = Session{
		ID:    "666",
		Email: "test@test.co.uk",
	}

	testHost = "http://localhost:8082"
	testErr  = errors.New("this is a test error")
)

func TestZebedeeClient_ListUserKeyring(t *testing.T) {
	Convey("Given httpCli.Do returns an error", t, func() {
		mockHttpCli := &mock.HttpClientMock{
			DoFunc: func(ctx context.Context, req *http.Request) (*http.Response, error) {
				return nil, testErr
			},
		}

		zebedeeCli := newZebedeeCli(mockHttpCli)

		Convey("When zebedeeCli.ListUserKeyring is called", func() {
			keyring, err := zebedeeCli.ListUserKeyring(testSess)

			Convey("Then the zebedee client is called with expected arguments", func() {
				assertHttpCliArguments(mockHttpCli)
			})

			Convey("And the expected error is returned", func() {
				So(err, ShouldNotBeNil)
				So(err, ShouldResemble, testErr)
				So(keyring, ShouldBeEmpty)
			})
		})
	})

	Convey("Given a non 200 response is returned", t, func() {
		mockHttpCli := &mock.HttpClientMock{
			DoFunc: func(ctx context.Context, req *http.Request) (*http.Response, error) {
				return &http.Response{
					StatusCode: 500,
					Request:    httptest.NewRequest(http.MethodGet, "/ListKeyring", nil),
					Body:       io.NopCloser(strings.NewReader("")),
				}, nil
			},
		}

		zebedeeCli := newZebedeeCli(mockHttpCli)

		Convey("When zebedeeCli.ListUserKeyring is called", func() {
			keyring, err := zebedeeCli.ListUserKeyring(testSess)

			Convey("Then the zebedee client is called with expected arguments", func() {
				assertHttpCliArguments(mockHttpCli)
			})

			Convey("And the expected response is returned", func() {
				So(keyring, ShouldBeNil)

				So(err, ShouldNotBeNil)

				var apiErr *APIError
				So(errors.As(err, &apiErr), ShouldBeTrue)
				So(apiErr.ActualStatus, ShouldEqual, http.StatusInternalServerError)
				So(apiErr.ExpectedStatus, ShouldEqual, http.StatusOK)
				So(apiErr.Message, ShouldEqual, fmt.Sprintf(incorrectStatusErrFmt, "GET", "/ListKeyring", 200, 500))
			})
		})
	})

	Convey("The API returns a successful response", t, func() {
		expected := []string{"one", "two", "three"}
		b, err := json.Marshal(expected)
		So(err, ShouldBeNil)

		mockHttpCli := &mock.HttpClientMock{
			DoFunc: func(ctx context.Context, req *http.Request) (*http.Response, error) {
				return &http.Response{
					StatusCode: http.StatusOK,
					Request:    httptest.NewRequest(http.MethodGet, "/ListKeyring", nil),
					Body:       io.NopCloser(bytes.NewReader(b)),
				}, nil
			},
		}

		zebedeeCli := newZebedeeCli(mockHttpCli)

		Convey("When zebedeeCli.ListUserKeyring is called", func() {
			keyring, err := zebedeeCli.ListUserKeyring(testSess)

			Convey("Then the zebedee client is called with expected arguments", func() {
				assertHttpCliArguments(mockHttpCli)
			})

			Convey("And the expected response is returned", func() {
				So(keyring, ShouldResemble, expected)
				So(err, ShouldBeNil)
			})
		})
	})
}

func assertHttpCliArguments(httpCli *mock.HttpClientMock) {
	So(httpCli.DoCalls(), ShouldHaveLength, 1)

	call := httpCli.DoCalls()[0]
	So(call.Req.URL.RequestURI(), ShouldEqual, "/ListKeyring")
	So(call.Req.Method, ShouldEqual, http.MethodGet)
	So(call.Req.Header.Get("content-type"), ShouldEqual, "application/json")
	So(call.Req.Header.Get(request.FlorenceHeaderKey), ShouldEqual, testSess.ID)
}

func newZebedeeCli(httpCli HttpClient) *zebedeeClient {
	return &zebedeeClient{
		Host:       testHost,
		HttpClient: httpCli,
	}
}
