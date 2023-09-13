package zebedee

import (
	"fmt"
	"io"
	"net/http"
)

func (z *zebedeeClient) GetContent(s Session, collectionName string, path string) ([]byte, error) {
	uri := fmt.Sprintf("/content/%s?uri=%s", collectionName, path)
	req, err := z.newAuthenticatedRequest(uri, s.ID, http.MethodGet, s.IsServiceToken, nil)
	if err != nil {
		return nil, err
	}

	resp, err := z.do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if err = checkResponseStatus(resp, http.StatusOK); err != nil {
		return nil, err
	}

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return b, nil
}
