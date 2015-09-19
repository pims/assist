package assist

import (
	"encoding/json"
)

type TeamsService struct {
	client *Client
}

func NewTeamsService(client *Client) *TeamsService {
	return &TeamsService{client: client}
}

func (s *TeamsService) Members(teamName string) ([]*User, error) {
	body, err := s.client.get("/teams/" + teamName + "/members")
	if err != nil {
		return nil, err
	}

	members := make([]*User, 0)
	jsonErr := json.Unmarshal(body, &members)
	return members, jsonErr
}

func (s *TeamsService) Shots(teamName string) ([]*Shot, error) {
	body, err := s.client.get("/teams/" + teamName + "/shots")
	if err != nil {
		return nil, err
	}

	shots := make([]*Shot, 0)
	jsonErr := json.Unmarshal(body, &shots)
	return shots, jsonErr
}
