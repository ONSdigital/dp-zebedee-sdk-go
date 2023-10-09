package zebedee

import (
	"fmt"
	"net/http"
)

// AddTeamMember add a CMS user to the specified team
func (z *zebedeeClient) AddTeamMember(s Session, teamName, email string) error {
	uri := fmt.Sprintf("/teams/%s?email=%s", teamName, email)
	req, err := z.newAuthenticatedRequest(uri, s.ID, http.MethodPost, nil)
	if err != nil {
		return err
	}

	return z.executeRequestNoResponse(req, http.StatusOK)
}

// RemoveTeamMember remove a user from the specific team
func (z *zebedeeClient) RemoveTeamMember(s Session, teamName, email string) error {
	uri := fmt.Sprintf("/teams/%s?email=%s", teamName, email)
	req, err := z.newAuthenticatedRequest(uri, s.ID, http.MethodDelete, nil)
	if err != nil {
		return err
	}

	return z.executeRequestNoResponse(req, http.StatusOK)
}

// CreateTeam create a new team
func (z *zebedeeClient) CreateTeam(s Session, teamName string) (bool, error) {
	uri := fmt.Sprintf("/teams/%s", teamName)
	req, err := z.newAuthenticatedRequest(uri, s.ID, http.MethodPost, nil)
	if err != nil {
		return false, err
	}

	var result bool
	if err := z.requestObject(req, http.StatusOK, &result); err != nil {
		return false, err
	}

	return result, nil
}

// DeleteTeam delete a team
func (z *zebedeeClient) DeleteTeam(s Session, teamName string) error {
	uri := fmt.Sprintf("/teams/%s", teamName)
	req, err := z.newAuthenticatedRequest(uri, s.ID, http.MethodDelete, nil)
	if err != nil {
		return err
	}

	return z.executeRequestNoResponse(req, http.StatusOK)
}

// ListTeams return a list of the current teams in the CMS
func (z *zebedeeClient) ListTeams(s Session) (TeamsList, error) {
	var teams TeamsList
	req, err := z.newAuthenticatedRequest("/teams", s.ID, http.MethodGet, nil)
	if err != nil {
		return teams, err
	}

	if err := z.requestObject(req, http.StatusOK, &teams); err != nil {
		return teams, err
	}

	return teams, nil
}

// GetTeam return the team with the specified name.
func (z *zebedeeClient) GetTeam(s Session, teamName string) (Team, error) {
	var team Team
	uri := fmt.Sprintf("/teams/%s", teamName)
	req, err := z.newAuthenticatedRequest(uri, s.ID, http.MethodGet, nil)
	if err != nil {
		return team, err
	}

	if err := z.requestObject(req, http.StatusOK, &team); err != nil {
		return team, err
	}

	return team, nil
}
