package zebedee

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/ONSdigital/dp-net/v2/request"
)

// CollectionsAPI defines the collections endpoints in Zebedee CMS
type CollectionsAPI interface {
	GetCollectionByID(s Session, id string) (CollectionDescription, error)
	CreateCollection(s Session, desc CollectionDescription) (CollectionDescription, error)
	DeleteCollection(s Session, id string) error
	GetCollections(s Session) ([]CollectionDescription, error)
	UpdateCollection(s Session, desc CollectionDescription) error
	UpdateCollectionContent(s Session, id, contentUri string, content interface{}) error
	DeleteCollectionContent(s Session, id, contentUri string) error
	CompleteCollectionContent(s Session, id string, contentUri string) error
	ReviewCollectionContent(s Session, id string, contentUri string) error
	ApproveCollection(s Session, id string) error
	UnlockCollection(s Session, id string) error
	PublishCollection(s Session, id string) error
	GetCollectionDetails(s Session, id string) (CollectionDetails, error)
}

// PermissionsAPI defines the user permissions endpoints in Zebedee CMS
type PermissionsAPI interface {
	SetPermissions(s Session, p Permissions) error
	GetPermissions(s Session, email string) (Permissions, error)
}

// UsersAPI defines the user endpoints in Zebedee CMS
type UsersAPI interface {
	CreateUser(s Session, u User) (User, error)
	GetUser(s Session, email string) (User, error)
	GetUsers(s Session) ([]User, error)
	DeleteUser(s Session, email string) error
	SetPassword(s Session, c Credentials) error
}

// AuthAPI defines the authentication endpoints in Zebedee CMS
type AuthAPI interface {
	OpenSession(c Credentials) (Session, error)
	OpenSessionJWT(authToken string) (Session, error)
}

// TeamsAPI defines the teams endpoints in Zebedee CMS
type TeamsAPI interface {
	AddTeamMember(s Session, teamName, email string) error
	RemoveTeamMember(s Session, teamName, email string) error
	CreateTeam(s Session, teamName string) (bool, error)
	DeleteTeam(s Session, teamName string) error
	ListTeams(s Session) (TeamsList, error)
	GetTeam(s Session, teamName string) (Team, error)
}

// KeyringAPI defines the Keyring endpoints in Zebedee CMS
type KeyringAPI interface {
	ListUserKeyring(s Session) ([]string, error)
}

type ContentAPI interface {
	GetContent(s Session, collectionName string, uri string) ([]byte, error)
}

// Client defines a client for the Zebedee CMS API
type Client interface {
	AuthAPI
	UsersAPI
	PermissionsAPI
	CollectionsAPI
	TeamsAPI
	KeyringAPI
	ContentAPI
}

type zebedeeClient struct {
	Host       string
	HttpClient HttpClient
}

// NewClient create a new Client
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
	req.Header.Set(request.FlorenceHeaderKey, authToken)

	return req, nil
}

// requestObject execute a JSON http request and unmarshal the response into the provided entity
func (z *zebedeeClient) requestObject(r *http.Request, expectedStatus int, entity interface{}) error {
	resp, err := z.do(r)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if err = checkResponseStatus(resp, expectedStatus); err != nil {
		return err
	}

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(b, &entity); err != nil {
		return err
	}

	return nil
}

func (z *zebedeeClient) do(req *http.Request) (*http.Response, error) {
	return z.HttpClient.Do(req.Context(), req)
}

// discardResponse consume the response body and send it to dev/null
func discardResponse(resp *http.Response) error {
	_, err := io.Copy(io.Discard, resp.Body)
	return err
}

// executeRequestNoResponse execute the HTTP request, check for the expected status but discard the response body
func (z *zebedeeClient) executeRequestNoResponse(r *http.Request, expectedStatus int) error {
	resp, err := z.do(r)
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
