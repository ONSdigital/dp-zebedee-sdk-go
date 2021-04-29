package zebedee

import (
	"fmt"
	"net/http"
)

//AddTeamMember add a CMS user to the specified team
func (z *zebedeeClient) AddTeamMember(s Session, teamName, email string) error {
	uri := fmt.Sprintf("/teams/%s?email=%s", teamName, email)
	req, err := z.newAuthenticatedRequest(uri, s.ID, http.MethodPost, nil)
	if err != nil {
		return err
	}

	resp, err := z.HttpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if err = checkResponseStatus(resp, http.StatusOK); err != nil {
		return err
	}

	if err = discardResponse(resp); err != nil {
		return err
	}

	return nil
}
