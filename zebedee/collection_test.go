package zebedee

import (
	"errors"
	"github.com/ONSdigital/dp-net/request"
	"github.com/ONSdigital/dp-zebedee-sdk-go/zebedee/mock"
	"net/http"
	"net/http/httptest"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

const (
	host = "http://localhost:8082"
)

func Test_CreateCollection(t *testing.T) {
	httpClient := newHttpClientMock()
	zebedeeClient := NewClient(host, httpClient)
	session := newSession()

	Convey("Given a description of a new collection to create", t, func() {
		collectionDescription := NewCollection("Collection Name")

		Convey("When CreateCollection is called", func() {
			_, err := zebedeeClient.CreateCollection(session, collectionDescription)

			Convey("Then the expected request is sent to the HTTP client", func() {
				So(httpClient.RequestObjectCalls(), ShouldHaveLength, 1)

				req := httpClient.RequestObjectCalls()[0].R
				So(req.Method, ShouldEqual, http.MethodPost)
				So(req.URL.String(), ShouldEqual, host+"/collection")
				So(req.Header.Get(request.FlorenceHeaderKey), ShouldEqual, session.ID)
				So(req.Header.Get("content-type"), ShouldEqual, "application/json")

				So(httpClient.RequestObjectCalls()[0].ExpectedStatus, ShouldEqual, http.StatusOK)
			})

			Convey("Then no error is returned", func() {
				So(err, ShouldBeNil)
			})
		})
	})
}

func Test_CreateCollection_RequestError(t *testing.T) {
	expectedError := errors.New("something broke")
	httpClient := newHttpClientMock()
	httpClient.RequestObjectFunc = func(r *http.Request, expectedStatus int, entity interface{}) error {
		return expectedError
	}
	zebedeeClient := NewClient(host, httpClient)
	session := newSession()

	Convey("Given an error is returned from request object", t, func() {
		collectionDescription := NewCollection("Collection Name")

		Convey("When CreateCollection is called", func() {
			_, err := zebedeeClient.CreateCollection(session, collectionDescription)

			Convey("Then the expected error is returned", func() {
				So(httpClient.RequestObjectCalls(), ShouldHaveLength, 1)
				So(err, ShouldEqual, expectedError)
			})
		})
	})
}

func newSession() Session {
	return Session{
		Email: "testuser@zebedeesdktest.com",
		ID:    "54345",
	}
}

func newHttpClientMock() *mock.HttpClientMock {
	return &mock.HttpClientMock{
		DoFunc: func(r *http.Request) (*http.Response, error) {
			return httptest.NewRecorder().Result(), nil
		},
		RequestObjectFunc: func(r *http.Request, expectedStatus int, entity interface{}) error {
			return nil
		},
	}
}
