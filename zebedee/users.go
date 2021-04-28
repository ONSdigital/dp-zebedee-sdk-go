package zebedee

import (
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
