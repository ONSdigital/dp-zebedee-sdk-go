package zebedee

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
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

// Permissions is the model representing user's CMS permissions
type Permissions struct {
	Email  string `json:"email"`
	Admin  bool   `json:"admin"`
	Editor bool   `json:"editor"`
}

// OpenSession opens a new user session using the login credentials provided
func OpenSession(cli Client, c Credentials) (Session, error) {
	var s Session
	body, err := json.Marshal(c)
	if err != nil {
		return s, err
	}

	url := fmt.Sprintf("%s/login", cli.GetHost())
	r, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(body))
	if err != nil {
		return s, err
	}

	resp, err := cli.Do(r)
	if err != nil {
		return s, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return s, IncorrectStatusErr("/login", http.MethodPost, http.StatusOK, resp.StatusCode)
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return s, err
	}

	s = Session{
		Email: c.Email,
		ID:    string(b),
	}

	return s, nil
}

// SetPermissions  set the user's CMS permissions
func SetPermissions(cli Client, s Session, p Permissions) error {
	r, err := cli.NewAuthenticatedRequest("/permission", s.ID, http.MethodPost, p)
	if err != nil {
		return err
	}

	resp, err := cli.Do(r)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return IncorrectStatusErr("/permission", http.MethodPost, http.StatusOK, resp.StatusCode)
	}

	io.Copy(ioutil.Discard, resp.Body)
	return nil
}

// GetPermissions  get the user's CMS permissions
func GetPermissions(cli Client, s Session, email string) (Permissions, error) {
	var p Permissions
	uri := fmt.Sprintf("/permisson?email=%s", email)
	r, err := cli.NewAuthenticatedRequest(uri, s.ID, http.MethodGet, nil)
	if err != nil {
		return p, err
	}

	resp, err := cli.Do(r)
	if err != nil {
		return p, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return p, IncorrectStatusErr("/permission", http.MethodGet, http.StatusOK, resp.StatusCode)
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return p, err
	}

	if err := json.Unmarshal(b, &p); err != nil {
		return p, err
	}

	return p, nil
}