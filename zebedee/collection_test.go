package zebedee

import (
	"bytes"
	"context"
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

	collectionJson := `{
		"name": "Coronavirus key indicators"
	}`

	httpClient := &mock.HttpClientMock{
		DoFunc: func(ctx context.Context, req *http.Request) (*http.Response, error) {
			recorder := httptest.NewRecorder()
			recorder.Code = http.StatusOK
			recorder.Body = bytes.NewBufferString(collectionJson)
			return recorder.Result(), nil
		},
	}
	zebedeeClient := NewClient(host, httpClient)
	session := newSession()

	Convey("Given a description of a new collection to create", t, func() {
		collectionDescription := NewCollection("Collection Name")

		Convey("When CreateCollection is called", func() {
			createdCollection, err := zebedeeClient.CreateCollection(session, collectionDescription)

			Convey("Then the expected request is sent to the HTTP client", func() {
				So(httpClient.DoCalls(), ShouldHaveLength, 1)

				req := httpClient.DoCalls()[0].Req
				So(req.Method, ShouldEqual, http.MethodPost)
				So(req.URL.String(), ShouldEqual, host+"/collection")
				So(req.Header.Get(request.FlorenceHeaderKey), ShouldEqual, session.ID)
				So(req.Header.Get("content-type"), ShouldEqual, "application/json")
			})

			Convey("Then no error is returned", func() {
				So(err, ShouldBeNil)
			})

			Convey("Then the returned collection is as returned from Zebedee", func() {
				So(createdCollection, ShouldNotBeNil)
				So(createdCollection.Name, ShouldEqual, "Coronavirus key indicators")
			})
		})
	})
}

func Test_CreateCollection_RequestError(t *testing.T) {
	expectedError := errors.New("something broke")
	httpClient := newHttpClientMock()
	httpClient.DoFunc = func(ctx context.Context, req *http.Request) (*http.Response, error) {
		return nil, expectedError
	}
	zebedeeClient := NewClient(host, httpClient)
	session := newSession()

	Convey("Given an error is returned from request object", t, func() {
		collectionDescription := NewCollection("Collection Name")

		Convey("When CreateCollection is called", func() {
			_, err := zebedeeClient.CreateCollection(session, collectionDescription)

			Convey("Then the expected error is returned", func() {
				So(httpClient.DoCalls(), ShouldHaveLength, 1)
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
		DoFunc: func(ctx context.Context, req *http.Request) (*http.Response, error) {
			return httptest.NewRecorder().Result(), nil
		},
	}
}
