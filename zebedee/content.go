package zebedee

import (
	"context"
	"fmt"
	"io"
	"net/http"

	"github.com/ONSdigital/log.go/v2/log"
)

func (z *zebedeeClient) GetContent(s Session, collectionName string, path string) ([]byte, error) {
	ctx := context.Background()
	uri := fmt.Sprintf("/content/%s?uri=%s", collectionName, path)
	req, err := z.newAuthenticatedRequest(uri, s.ID, http.MethodGet, nil)
	if err != nil {
		return nil, err
	}

	resp, err := z.do(req)
	if err != nil {
		return nil, err
	}

	defer func() {
		if err := resp.Body.Close(); err != nil {
			log.Error(ctx, "error closing http response body", err)
		}
	}()

	if err = checkResponseStatus(resp, http.StatusOK); err != nil {
		return nil, err
	}

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return b, nil
}
