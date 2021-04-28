package zebedee

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

//Client defines a client for Zebedee CMS API
type Client interface {
	OpenSession(c Credentials) (Session, error)
	SetPermissions(s Session, p Permissions) error
	GetPermissions(s Session, email string) (Permissions, error)
	CreateUser(s Session, u User) (User, error)
	GetUsers(s Session) ([]User, error)
}

type zebedeeClient struct {
	Host       string
	HttpClient HttpClient
}

//NewClient create a new Client
func NewClient(host string, httpCli HttpClient) Client {
	return &zebedeeClient{
		Host:       host,
		HttpClient: httpCli,
	}
}

func (z *zebedeeClient) newAuthenticatedRequest(uri, authToken, method string, entity interface{}) (*http.Request, error) {
	var body io.Reader
	if entity != nil {
		b, err := json.Marshal(entity)
		if err != nil {
			return nil, err
		}

		body = bytes.NewReader(b)
	}

	url := fmt.Sprintf("%s%s", z.Host, uri)
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}

	req.Header.Set("content-type", "application/json")
	req.Header.Set("X-Florence-Token", authToken)
	return req, nil
}
