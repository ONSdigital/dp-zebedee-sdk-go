package zebedee

import (
	"io"
	"io/ioutil"
	"net/http"
)

//CreateUser a new CMS user
func (z *zebedeeClient) CreateUser(s Session, u User) (User, error) {
	var user User
	req, err := z.newAuthenticatedRequest("/users", s.ID, http.MethodPost, u)
	if err != nil {
		return user, err
	}

	err = z.HttpClient.RequestObject(req, http.StatusOK, &user)
	if err != nil {
		return user, err
	}
	return user, nil
}

//GetUsers a list of the CMS users
func (z *zebedeeClient) GetUsers(s Session) ([]User, error) {
	req, err := z.newAuthenticatedRequest("/users", s.ID, http.MethodGet, nil)
	if err != nil {
		return nil, err
	}

	var users []User
	err = z.HttpClient.RequestObject(req, http.StatusOK, &users)
	if err != nil {
		return nil, err
	}

	return users, nil
}

//DeleteUser delete a CMS user.
func (z *zebedeeClient) DeleteUser(s Session, email string) error {
	req, err := z.newAuthenticatedRequest("/users?email="+email, s.ID, http.MethodDelete, nil)
	if err != nil {
		return err
	}

	resp, err := z.HttpClient.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return IncorrectStatusErr(req, http.StatusOK, resp.StatusCode)
	}

	_, err = io.Copy(ioutil.Discard, resp.Body)
	if err != nil {
		return err
	}

	return nil
}
