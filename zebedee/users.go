package zebedee

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type User struct {
	Name              string `json:"name"`
	Email             string `json:"email"`
	Inactive          bool   `json:"inactive"`
	LastAdmin         string `json:"lastAdmin"`
	TemporaryPassword bool   `json:"temporaryPassword"`
}

//CreateUser a new CMS user
func CreateUser(cli Client, s Session, u User) (User, error) {
	var user User
	req, err := cli.NewAuthenticatedRequest("/users", s.ID, http.MethodPost, u)
	if err != nil {
		return user, err
	}

	resp, err := cli.Do(req)
	if err != nil {
		return user, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return user, fmt.Errorf("error creating user expected status 200 but was %d", resp.StatusCode)
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return user, err
	}

	if err := json.Unmarshal(b, &user); err != nil {
		return user, err
	}

	return user, nil
}

//GetUsers a list of the CMS users
func GetUsers(cli Client, s Session) ([]User, error) {
	req, err := cli.NewAuthenticatedRequest("/users", s.ID, http.MethodGet, nil)
	if err != nil {
		return nil, err
	}

	resp, err := cli.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error creating user expected status 200 but was %d", resp.StatusCode)
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var users []User
	if err := json.Unmarshal(b, &users); err != nil {
		return nil, err
	}

	return users, nil
}

func GetUsers2(cli Client, s Session) ([]User, error) {
	req, err := cli.NewAuthenticatedRequest("/users", s.ID, http.MethodGet, nil)
	if err != nil {
		return nil, err
	}

	var users []User
	err = cli.RequestObject(req, http.StatusOK, &users)
	if err != nil {
		return nil, err
	}

	return users, nil
}
