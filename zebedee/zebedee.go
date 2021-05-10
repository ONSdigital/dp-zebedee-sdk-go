package zebedee

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

//CollectionsAPI defines the collections endpoints in Zebedee CMS
type CollectionsAPI interface {
	GetCollectionByID(s Session, id string) (CollectionDescription, error)
	CreateCollection(s Session, desc CollectionDescription) (CollectionDescription, error)
	DeleteCollection(s Session, id string) error
}

//PermissionsAPI defines the user permissions endpoints in Zebedee CMS
type PermissionsAPI interface {
	SetPermissions(s Session, p Permissions) error
	GetPermissions(s Session, email string) (Permissions, error)
}

//UsersAPI defines the user endpoints in Zebedee CMS
type UsersAPI interface {
	CreateUser(s Session, u User) (User, error)
	GetUser(s Session, email string) (User, error)
	GetUsers(s Session) ([]User, error)
	DeleteUser(s Session, email string) error
	SetPassword(s Session, email, password string) error
}

//AuthAPI defines the authentication endpoints in Zebedee CMS
type AuthAPI interface {
	OpenSession(c Credentials) (Session, error)
}

//TeamsAPI defines the teams endpoints in Zebedee CMS
type TeamsAPI interface {
	AddTeamMember(s Session, teamName, email string) error
}

//Client defines a client for the Zebedee CMS API
type Client interface {
	AuthAPI
	UsersAPI
	PermissionsAPI
	CollectionsAPI
	TeamsAPI
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
