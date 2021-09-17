package zebedee

import (
	"fmt"
	"net/http"
)

//ListUserKeyring returns a list of collection ID's the keys the user has stored on their keyring
func (z *zebedeeClient) ListUserKeyring(s Session, email, src string) ([]string, error) {
	uri := fmt.Sprintf("/ListKeyring?email=%s&src=%s", email, src)
	req, err := z.newAuthenticatedRequest(uri, s.ID, http.MethodGet, nil)
	if err != nil {
		return nil, err
	}

	var keys []string
	err = z.requestObject(req, 200, &keys)
	if err != nil {
		return nil, err
	}

	return keys, nil
}
