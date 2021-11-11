package zebedee

import (
	"fmt"
	"net/http"
)

//GetCollectionDetails return the requested collection details from the CMS.
func (z *zebedeeClient) GetCollectionDetails(s Session, id string) (CollectionDetails, error) {
	var details CollectionDetails

	uri := fmt.Sprintf("/collectionDetails/%s", id)
	req, err := z.newAuthenticatedRequest(uri, s.ID, http.MethodGet, nil)
	if err != nil {
		return details, err
	}

	err = z.requestObject(req, 200, &details)
	if err != nil {
		return details, err
	}

	return details, nil
}
