package zebedee

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"github.com/ONSdigital/dp-net/request"
	"github.com/ONSdigital/dp-zebedee-sdk-go/zebedee/mock"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

const (
	host         = "http://localhost:8082"
	uri          = "/the/uri"
	collectionId = "collectionID"
	pageContent  = `{"type":"static_page"}`
)

func Test_CreateCollection(t *testing.T) {

	responseBody := `{
		"name": "Coronavirus key indicators"
	}`
	httpClient := mockHttpResponse(http.StatusOK, responseBody)
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

func Test_CreateCollection_HttpError(t *testing.T) {
	expectedError := errors.New("something broke")
	session := newSession()
	collectionDescription := NewCollection("Collection Name")

	Convey("Given an error is returned from the HTTP client", t, func() {
		httpClient := mockHttpError(expectedError)
		zebedeeClient := NewClient(host, httpClient)

		Convey("When CreateCollection is called", func() {
			_, err := zebedeeClient.CreateCollection(session, collectionDescription)

			Convey("Then the expected error is returned", func() {
				So(httpClient.DoCalls(), ShouldHaveLength, 1)
				So(err, ShouldEqual, expectedError)
			})
		})
	})
}

func Test_DeleteCollection(t *testing.T) {
	responseBody := `true`
	httpClient := mockHttpResponse(http.StatusOK, responseBody)
	zebedeeClient := NewClient(host, httpClient)
	session := newSession()

	Convey("Given an ID of a collection to delete", t, func() {
		collectionId := "1234"

		Convey("When DeleteCollection is called", func() {
			err := zebedeeClient.DeleteCollection(session, collectionId)

			Convey("Then the expected request is sent to the HTTP client", func() {
				So(httpClient.DoCalls(), ShouldHaveLength, 1)

				req := httpClient.DoCalls()[0].Req
				So(req.Method, ShouldEqual, http.MethodDelete)
				So(req.URL.String(), ShouldEqual, host+"/collection/"+collectionId)
				So(req.Header.Get(request.FlorenceHeaderKey), ShouldEqual, session.ID)
				So(req.Header.Get("content-type"), ShouldEqual, "application/json")
			})

			Convey("Then no error is returned", func() {
				So(err, ShouldBeNil)
			})
		})
	})
}

func Test_DeleteCollection_FalseResponse(t *testing.T) {
	session := newSession()

	Convey("Given an false response body is returned from Zebedee", t, func() {
		responseBody := `false`
		httpClient := mockHttpResponse(http.StatusOK, responseBody)
		zebedeeClient := NewClient(host, httpClient)

		Convey("When DeleteCollection is called", func() {
			err := zebedeeClient.DeleteCollection(session, collectionId)

			Convey("Then the expected error is returned", func() {
				So(httpClient.DoCalls(), ShouldHaveLength, 1)
				So(err, ShouldNotBeNil)
				So(err.Error(), ShouldEqual, "delete collection request unsuccessful: "+collectionId)
			})
		})
	})
}

func Test_DeleteCollection_HttpError(t *testing.T) {
	session := newSession()

	Convey("Given an error is returned from the HTTP client", t, func() {
		expectedError := errors.New("something broke")
		httpClient := mockHttpError(expectedError)
		zebedeeClient := NewClient(host, httpClient)

		Convey("When DeleteCollection is called", func() {
			err := zebedeeClient.DeleteCollection(session, collectionId)

			Convey("Then the expected error is returned", func() {
				So(httpClient.DoCalls(), ShouldHaveLength, 1)
				So(err, ShouldEqual, expectedError)
			})
		})
	})
}

func Test_UpdateCollectionContent(t *testing.T) {
	responseBody := `true`
	httpClient := mockHttpResponse(http.StatusOK, responseBody)
	zebedeeClient := NewClient(host, httpClient)
	session := newSession()

	Convey("Given a request to update collection content", t, func() {
		content := getContent()

		Convey("When UpdateCollectionContent is called", func() {
			err := zebedeeClient.UpdateCollectionContent(session, collectionId, uri, content)

			Convey("Then the expected request is sent to the HTTP client", func() {
				So(httpClient.DoCalls(), ShouldHaveLength, 1)

				req := httpClient.DoCalls()[0].Req
				So(req.Method, ShouldEqual, http.MethodPost)
				So(req.URL.String(), ShouldEqual, host+"/content/collectionID?uri=/the/uri&overwriteExisting=true&recursive=false&validateJson=true")
				So(req.Header.Get(request.FlorenceHeaderKey), ShouldEqual, session.ID)
				So(req.Header.Get("content-type"), ShouldEqual, "application/json")

				bodyBytes, _ := ioutil.ReadAll(req.Body)
				bodyContent := string(bodyBytes)
				So(bodyContent, ShouldEqual, pageContent)
			})

			Convey("Then no error is returned", func() {
				So(err, ShouldBeNil)
			})
		})
	})
}

func Test_UpdateCollectionContent_FalseResponse(t *testing.T) {
	session := newSession()

	Convey("Given a mocked zebedee response that returns false", t, func() {
		responseBody := `false`
		httpClient := mockHttpResponse(http.StatusOK, responseBody)
		zebedeeClient := NewClient(host, httpClient)
		content := getContent()

		Convey("When UpdateCollectionContent is called", func() {
			err := zebedeeClient.UpdateCollectionContent(session, collectionId, uri, content)

			Convey("Then the expected error is returned", func() {
				So(httpClient.DoCalls(), ShouldHaveLength, 1)
				So(err, ShouldNotBeNil)
				So(err.Error(), ShouldEqual, "update collection content request unsuccessful: "+collectionId)
			})
		})
	})
}

func Test_UpdateCollectionContent_HttpError(t *testing.T) {
	session := newSession()

	Convey("Given an error is returned from the HTTP client", t, func() {
		expectedError := errors.New("something broke")
		httpClient := mockHttpError(expectedError)
		zebedeeClient := NewClient(host, httpClient)
		content := getContent()

		Convey("When UpdateCollectionContent is called", func() {
			err := zebedeeClient.UpdateCollectionContent(session, collectionId, uri, content)

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

func mockHttpError(err error) *mock.HttpClientMock {
	return &mock.HttpClientMock{
		DoFunc: func(ctx context.Context, req *http.Request) (*http.Response, error) {
			return nil, err
		},
	}
}

func mockHttpResponse(responseCode int, responseBody string) *mock.HttpClientMock {
	return &mock.HttpClientMock{
		DoFunc: func(ctx context.Context, req *http.Request) (*http.Response, error) {
			recorder := httptest.NewRecorder()
			recorder.Code = responseCode
			recorder.Body = bytes.NewBufferString(responseBody)
			return recorder.Result(), nil
		},
	}
}

func getContent() interface{} {
	var content interface{}
	err := json.Unmarshal([]byte(pageContent), &content)
	So(err, ShouldBeNil)
	return content
}
