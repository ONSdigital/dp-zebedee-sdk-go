package auth

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/ONSdigital/dp-zebedee-sdk-go/zhttp"
)

const (
	loginStatusErr = "login returned incorrect status code expected 200 but was %d"
)

// Credentials is the model representing the user login details
type Credentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Session is the model of a CMS user session.
type Session struct {
	Email string `json:"email"`
	ID    string `json:"id"`
}

// OpenSession opens a new user session using the login credentials provided
func OpenSession(cli zhttp.Client, host string, c Credentials) (*Session, error) {
	body, err := json.Marshal(c)
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("%s/login", host)
	r, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	resp, err := cli.Do(r)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(loginStatusErr, resp.StatusCode)
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	sess := &Session{
		Email: c.Email,
		ID:    string(b),
	}

	return sess, nil
}

// {name: "", email: ""} users
// {email: "", password: ""}  password
// {email: "", admin: false, editor: false} permission
