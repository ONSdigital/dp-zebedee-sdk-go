package zebedee

import (
	"net/http"
)

// ListUserKeyring returns a list of collection ID's for the keys the user has access to.
func (z *zebedeeClient) ListUserKeyring(s Session) ([]string, error) {
	uri := "/ListKeyring"
	req, err := z.newAuthenticatedRequest(uri, s.ID, http.MethodGet, s.IsServiceToken, nil)
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
