package assist

import (
	"encoding/json"
)

type TeamsService struct {
	client *Client
}

// NewTeamsService creates a new TeamService client
func NewTeamsService(client *Client) *TeamsService {
	return &TeamsService{client: client}
}

// Retrieves list of team members for the given team
func (s *TeamsService) Members(teamName string) ([]*User, error) {
	body, err := s.client.get("/teams/" + teamName + "/members")
	if err != nil {
		return nil, err
	}

	members := make([]*User, 0)
	jsonErr := json.Unmarshal(body, &members)
	return members, jsonErr
}

// Retrieves list of team shots for the given team
func (s *TeamsService) Shots(teamName string) ([]*Shot, error) {
	body, err := s.client.get("/teams/" + teamName + "/shots")
	if err != nil {
		return nil, err
	}

	shots := make([]*Shot, 0)
	jsonErr := json.Unmarshal(body, &shots)
	return shots, jsonErr
}
