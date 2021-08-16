// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package mock

import (
	"net/http"
	"sync"
)

// HttpClientMock is a mock implementation of zebedee.HttpClient.
//
//     func TestSomethingThatUsesHttpClient(t *testing.T) {
//
//         // make and configure a mocked zebedee.HttpClient
//         mockedHttpClient := &HttpClientMock{
//             DoFunc: func(r *http.Request) (*http.Response, error) {
// 	               panic("mock out the Do method")
//             },
//             RequestObjectFunc: func(r *http.Request, expectedStatus int, entity interface{}) error {
// 	               panic("mock out the RequestObject method")
//             },
//         }
//
//         // use mockedHttpClient in code that requires zebedee.HttpClient
//         // and then make assertions.
//
//     }
type HttpClientMock struct {
	// DoFunc mocks the Do method.
	DoFunc func(r *http.Request) (*http.Response, error)

	// RequestObjectFunc mocks the RequestObject method.
	RequestObjectFunc func(r *http.Request, expectedStatus int, entity interface{}) error

	// calls tracks calls to the methods.
	calls struct {
		// Do holds details about calls to the Do method.
		Do []struct {
			// R is the r argument value.
			R *http.Request
		}
		// RequestObject holds details about calls to the RequestObject method.
		RequestObject []struct {
			// R is the r argument value.
			R *http.Request
			// ExpectedStatus is the expectedStatus argument value.
			ExpectedStatus int
			// Entity is the entity argument value.
			Entity interface{}
		}
	}
	lockDo            sync.RWMutex
	lockRequestObject sync.RWMutex
}

// Do calls DoFunc.
func (mock *HttpClientMock) Do(r *http.Request) (*http.Response, error) {
	if mock.DoFunc == nil {
		panic("HttpClientMock.DoFunc: method is nil but HttpClient.Do was just called")
	}
	callInfo := struct {
		R *http.Request
	}{
		R: r,
	}
	mock.lockDo.Lock()
	mock.calls.Do = append(mock.calls.Do, callInfo)
	mock.lockDo.Unlock()
	return mock.DoFunc(r)
}

// DoCalls gets all the calls that were made to Do.
// Check the length with:
//     len(mockedHttpClient.DoCalls())
func (mock *HttpClientMock) DoCalls() []struct {
	R *http.Request
} {
	var calls []struct {
		R *http.Request
	}
	mock.lockDo.RLock()
	calls = mock.calls.Do
	mock.lockDo.RUnlock()
	return calls
}

// RequestObject calls RequestObjectFunc.
func (mock *HttpClientMock) RequestObject(r *http.Request, expectedStatus int, entity interface{}) error {
	if mock.RequestObjectFunc == nil {
		panic("HttpClientMock.RequestObjectFunc: method is nil but HttpClient.RequestObject was just called")
	}
	callInfo := struct {
		R              *http.Request
		ExpectedStatus int
		Entity         interface{}
	}{
		R:              r,
		ExpectedStatus: expectedStatus,
		Entity:         entity,
	}
	mock.lockRequestObject.Lock()
	mock.calls.RequestObject = append(mock.calls.RequestObject, callInfo)
	mock.lockRequestObject.Unlock()
	return mock.RequestObjectFunc(r, expectedStatus, entity)
}

// RequestObjectCalls gets all the calls that were made to RequestObject.
// Check the length with:
//     len(mockedHttpClient.RequestObjectCalls())
func (mock *HttpClientMock) RequestObjectCalls() []struct {
	R              *http.Request
	ExpectedStatus int
	Entity         interface{}
} {
	var calls []struct {
		R              *http.Request
		ExpectedStatus int
		Entity         interface{}
	}
	mock.lockRequestObject.RLock()
	calls = mock.calls.RequestObject
	mock.lockRequestObject.RUnlock()
	return calls
}
