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
	GetCollections(s Session) ([]CollectionDescription, error)
	UpdateCollection(s Session, desc CollectionDescription) error
	UpdateCollectionContent(s Session, id, contentUri string, content io.Reader, overwriteExisting, recursive, validateJson bool) error
	DeleteCollectionContent(s Session, id, contentUri string) error
	CompleteCollectionContent(s Session, id string, contentUri string, recursive bool) error
	ReviewCollectionContent(s Session, id string, contentUri string, recursive bool) error
	ApproveCollection(s Session, id string) error
	UnlockCollection(s Session, id string) error
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
	SetPassword(s Session, c Credentials) error
}

//AuthAPI defines the authentication endpoints in Zebedee CMS
type AuthAPI interface {
	OpenSession(c Credentials) (Session, error)
}

//TeamsAPI defines the teams endpoints in Zebedee CMS
type TeamsAPI interface {
	AddTeamMember(s Session, teamName, email string) error
	RemoveTeamMember(s Session, teamName, email string) error
	CreateTeam(s Session, teamName string) (bool, error)
	DeleteTeam(s Session, teamName string) error
	ListTeams(s Session) (TeamsList, error)
	GetTeam(s Session, teamName string) (Team, error)
}

//KeyringAPI defines the Keyring endpoints in Zebedee CMS
type KeyringAPI interface {
	ListUserKeyring(s Session, email string) ([]string, error)
}

//Client defines a client for the Zebedee CMS API
type Client interface {
	AuthAPI
	UsersAPI
	PermissionsAPI
	CollectionsAPI
	TeamsAPI
	KeyringAPI
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

//executeRequestNoResponse execute the HTTP request, check for the expected status but discard the response body
func (z *zebedeeClient) executeRequestNoResponse(r *http.Request, expectedStatus int) error {
	resp, err := z.HttpClient.Do(r)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if err = checkResponseStatus(resp, expectedStatus); err != nil {
		return err
	}

	if err = discardResponse(resp); err != nil {
		return err
	}

	return nil
}
