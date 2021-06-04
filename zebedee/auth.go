package zebedee

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// OpenSession opens a new user session using the login credentials provided
func (z *zebedeeClient) OpenSession(c Credentials) (Session, error) {
	var s Session
	body, err := json.Marshal(c)
	if err != nil {
		return s, err
	}

	url := fmt.Sprintf("%s/login", z.Host)
	r, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(body))
	if err != nil {
		return s, err
	}

	resp, err := z.HttpClient.Do(r)
	if err != nil {
		return s, err
	}
	defer resp.Body.Close()

	if err = checkResponseStatus(resp, http.StatusOK); err != nil {
		return s, err
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
func (z *zebedeeClient) SetPermissions(s Session, p Permissions) error {
	r, err := z.newAuthenticatedRequest("/permission", s.ID, http.MethodPost, p)
	if err != nil {
		return err
	}

	return z.executeRequestNoResponse(r, http.StatusOK)
}

// GetPermissions  get the user's CMS permissions
func (z *zebedeeClient) GetPermissions(s Session, email string) (Permissions, error) {
	var p Permissions
	uri := fmt.Sprintf("/permission?email=%s", email)

	r, err := z.newAuthenticatedRequest(uri, s.ID, http.MethodGet, nil)
	if err != nil {
		return p, err
	}

	err = z.HttpClient.RequestObject(r, http.StatusOK, &p)
	if err != nil {
		return p, err
	}

	return p, nil
}

//SetPassword set the user password
func (z *zebedeeClient) SetPassword(s Session, email, password string) error {
	c := Credentials{
		Email:    email,
		Password: password,
	}

	r, err := z.newAuthenticatedRequest("/password", s.ID, http.MethodPost, c)
	if err != nil {
		return err
	}

	return z.executeRequestNoResponse(r, http.StatusOK)

	return nil
}
