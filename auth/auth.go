package auth

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type Credentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Session struct {
	Email string
	ID    string
}

const (
	loginStatusErr = "login returned incorrect status code expected 200 but was %d"
)

var (
	httpCli = http.Client{
		Timeout: 5 * time.Second,
	}
)

func OpenSession(host string, c Credentials) (*Session, error) {
	body, err := json.Marshal(c)
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("%s/login", host)
	r, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	resp, err := httpCli.Do(r)
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
