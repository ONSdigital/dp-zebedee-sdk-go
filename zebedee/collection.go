package zebedee

import (
	"fmt"
	"io"
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

// ApproveCollection approves a collection with the provided ID.
// The approval can only take place once all collection content is reviewed
// A scheduled collection will only be published if the collection is approved
func (z *zebedeeClient) ApproveCollection(s Session, id string) error {
	uri := fmt.Sprintf("/approve/%s", id)
	req, err := z.newAuthenticatedRequest(uri, s.ID, http.MethodPost, nil)
	if err != nil {
		return err
	}

	var success bool
	err = z.HttpClient.RequestObject(req, 200, &success)
	if err != nil {
		return err
	}

	if !success {
		return fmt.Errorf("approve collection request unsuccesseful: %s", id)
	}

	return nil
}

// UnlockCollection reverses the approval state, allowing collection content to be edited
func (z *zebedeeClient) UnlockCollection(s Session, id string) error {
	uri := fmt.Sprintf("/unlock/%s", id)
	req, err := z.newAuthenticatedRequest(uri, s.ID, http.MethodPost, nil)
	if err != nil {
		return err
	}

	var success bool
	err = z.HttpClient.RequestObject(req, 200, &success)
	if err != nil {
		return err
	}

	if !success {
		return fmt.Errorf("unlock collection request unsuccesseful: %s", id)
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

//UpdateCollection updates the collection description
func (z *zebedeeClient) UpdateCollection(s Session, desc CollectionDescription) error {
	uri := fmt.Sprintf("/collection/%s", desc.ID)
	req, err := z.newAuthenticatedRequest(uri, s.ID, http.MethodPut, desc)
	if err != nil {
		return err
	}

	return z.executeRequestNoResponse(req, http.StatusOK)
}

//UpdateCollectionContent updates content within a collection
//  overwriteExisting (default:true) - if set to false, any existing content will not be overwritten
//  recursive (default:false) - if set to true, all associated files alongside the page will be added to the collection's in progress directory
//              if set to false, only the data.json file will be added to the collection's in progress directory
//  validateJson (default:true) - if set to true, the json will be validated to ensure it's a valid page JSON structure
func (z *zebedeeClient) UpdateCollectionContent(
	s Session,
	id, contentUri string,
	content io.Reader,
	overwriteExisting, recursive, validateJson bool) error {

	url := fmt.Sprintf("%s/content/%s?uri=%s&overwriteExisting=%t&recursive=%t&validateJson=%t",
		z.Host, id, contentUri, overwriteExisting, recursive, validateJson)

	req, err := http.NewRequest(http.MethodPost, url, content)
	if err != nil {
		return err
	}

	req.Header.Set("content-type", "application/json")
	req.Header.Set("X-Florence-Token", s.ID)

	return z.executeRequestNoResponse(req, http.StatusOK)
}

//DeleteCollectionContent deletes content from a collection
func (z *zebedeeClient) DeleteCollectionContent(s Session, id, contentUri string) error {
	url := fmt.Sprintf("%s/content/%s?uri=%s", z.Host, id, contentUri)

	req, err := http.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		return err
	}

	req.Header.Set("X-Florence-Token", s.ID)
	return z.executeRequestNoResponse(req, http.StatusOK)
}

//CompleteCollectionContent sets content in a collection to the complete state.
// This is done once the content has been updated and the user is satisfied that the changes are complete
func (z *zebedeeClient) CompleteCollectionContent(s Session, id, contentUri string, recursive bool) error {
	url := fmt.Sprintf("%s/complete/%s?uri=%s&recursive=%t", z.Host, id, contentUri, recursive)

	req, err := http.NewRequest(http.MethodPost, url, nil)
	if err != nil {
		return err
	}

	req.Header.Set("X-Florence-Token", s.ID)
	return z.executeRequestNoResponse(req, http.StatusOK)
}

//ReviewCollectionContent sets content in a collection to the reviewed state.
// This is done once the content has been reviewed by a user who is not the original editor.
func (z *zebedeeClient) ReviewCollectionContent(s Session, id, contentUri string, recursive bool) error {
	url := fmt.Sprintf("%s/review/%s?uri=%s&recursive=%t", z.Host, id, contentUri, recursive)

	req, err := http.NewRequest(http.MethodPost, url, nil)
	if err != nil {
		return err
	}

	req.Header.Set("X-Florence-Token", s.ID)
	return z.executeRequestNoResponse(req, http.StatusOK)
}
