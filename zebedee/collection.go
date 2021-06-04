package zebedee

import (
	"fmt"
	"net/http"
	"time"
)

//NewCollection create a new collection description using the default values.
func NewCollection(name string) CollectionDescription {
	return CollectionDescription{
		collectionBase: collectionBase{
			Name:        name,
			Type:        Manual,
			PublishDate: time.Now().Format(CollectionDateFMT),
			Teams:       make([]string, 0),
		},
		Encrypted:             false,
		PublishComplete:       false,
		ApprovalStatus:        "",
		InProgressUris:        nil,
		CompleteUris:          nil,
		ReviewedUris:          nil,
		DatasetVersions:       nil,
		Datasets:              nil,
		EventsByUri:           nil,
		PendingDeletes:        nil,
		PublishResults:        nil,
		TimeseriesImportFiles: nil,
	}
}

//CreateCollection create a new collection. Returns an updated collection description containing the generated collection ID or an error.
func (z *zebedeeClient) CreateCollection(s Session, desc CollectionDescription) (CollectionDescription, error) {
	var updated CollectionDescription

	uri := fmt.Sprintf("/collection")
	req, err := z.newAuthenticatedRequest(uri, s.ID, http.MethodPost, desc)
	if err != nil {
		return updated, err
	}

	err = z.HttpClient.RequestObject(req, http.StatusOK, &updated)
	if err != nil {
		return updated, err
	}

	return updated, nil
}

//GetCollectionByID get a collection by ID. Returns the collection description or an error.
func (z *zebedeeClient) GetCollectionByID(s Session, id string) (CollectionDescription, error) {
	var desc CollectionDescription

	uri := fmt.Sprintf("/collection/%s", id)
	req, err := z.newAuthenticatedRequest(uri, s.ID, http.MethodGet, nil)
	if err != nil {
		return desc, err
	}

	err = z.HttpClient.RequestObject(req, 200, &desc)
	if err != nil {
		return desc, err
	}

	return desc, nil
}

//DeleteCollection deletes a collection with the provided ID. Returns error if unsuccessful
func (z *zebedeeClient) DeleteCollection(s Session, id string) error {
	uri := fmt.Sprintf("/collection/%s", id)
	req, err := z.newAuthenticatedRequest(uri, s.ID, http.MethodDelete, nil)
	if err != nil {
		return err
	}

	var success bool
	err = z.HttpClient.RequestObject(req, 200, &success)
	if err != nil {
		return err
	}

	if !success {
		return fmt.Errorf("delete collection request unsuccesseful: %s", id)
	}

	return nil
}

//GetCollections returns a list of collection descriptions for each current collection
func (z *zebedeeClient) GetCollections(s Session) ([]CollectionDescription, error) {
	req, err := z.newAuthenticatedRequest("/collections", s.ID, http.MethodGet, nil)
	if err != nil {
		return nil, err
	}

	var collectionList []CollectionDescription
	err = z.HttpClient.RequestObject(req, 200, &collectionList)
	if err != nil {
		return nil, err
	}

	return collectionList, nil
}
