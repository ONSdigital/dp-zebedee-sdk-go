package users

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/ONSdigital/dp-zebedee-sdk-go/auth"
	"github.com/ONSdigital/dp-zebedee-sdk-go/zhttp"
)

type Model struct {
	Name              string `json:"name"`
	Email             string `json:"email"`
	Inactive          bool   `json:"inactive"`
	LastAdmin         string `json:"lastAdmin"`
	TemporaryPassword bool   `json:"temporaryPassword"`
}

//Create a new CMS user
func Create(cli zhttp.Client, host string, s auth.Session, u Model) (Model, error) {
	var user Model
	url := fmt.Sprintf("%s/users", host)
	req, err := zhttp.NewAuthenticatedRequest(url, s.ID, http.MethodPost, u)
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

//Get a list of the CMS users
func Get(cli zhttp.Client, host string, s auth.Session) ([]Model, error) {
	url := fmt.Sprintf("%s/users", host)
	req, err := zhttp.NewAuthenticatedRequest(url, s.ID, http.MethodGet, nil)
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

	var users []Model
	if err := json.Unmarshal(b, &users); err != nil {
		return nil, err
	}

	return users, nil
}
