package zebedee

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/ONSdigital/dp-net/request"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_GetContent(t *testing.T) {
	session := newSession()
	content := "content test"

	Convey("Given a mock HTTP client that returns a successful response", t, func() {
		httpClient := mockHttpResponse(http.StatusOK, content)
		zebedeeClient := NewClient(host, httpClient)
		expectedUrl := fmt.Sprintf("%s/content/%s?uri=%s", host, collectionId, uri)

		Convey("When GetContent is called", func() {
			res, err := zebedeeClient.GetContent(session, collectionId, uri)

			Convey("Then the expected request is sent to the HTTP client", func() {
				So(httpClient.DoCalls(), ShouldHaveLength, 1)

				req := httpClient.DoCalls()[0].Req
				So(req.Method, ShouldEqual, http.MethodGet)
				So(req.URL.String(), ShouldEqual, expectedUrl)
				So(req.Header.Get(request.FlorenceHeaderKey), ShouldEqual, session.ID)

				Convey("Then no error is returned", func() {
					So(err, ShouldBeNil)
					So(string(res), ShouldEqual, content)
				})
			})
		})
	})
}

func Test_GetContentError(t *testing.T) {
	session := newSession()

	Convey("Given a mock  HTTP client that returns an error response", t, func() {
		httpClient := mockHttpResponse(http.StatusBadGateway, "")
		zebedeeClient := NewClient(host, httpClient)
		expectedUrl := fmt.Sprintf("%s/content/%s?uri=%s", host, collectionId, uri)

		Convey("When GetContent is called", func() {
			res, err := zebedeeClient.GetContent(session, collectionId, uri)

			Convey("Then the expected request is sent to the HTTP client", func() {
				So(httpClient.DoCalls(), ShouldHaveLength, 1)

				req := httpClient.DoCalls()[0].Req
				So(req.Method, ShouldEqual, http.MethodGet)
				So(req.URL.String(), ShouldEqual, expectedUrl)
				So(req.Header.Get(request.FlorenceHeaderKey), ShouldEqual, session.ID)

				Convey("Then an error is returned", func() {
					So(err, ShouldNotBeNil)
					So(res, ShouldBeNil)
				})
			})
		})
	})
}
