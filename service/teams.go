package service

import (
	"encoding/json"
	"github.com/pims/assist"
)

type Teams struct {
	client *Client
}

// NewTeams creates a new TeamService client
func NewTeams(client *Client) *Teams {
	return &Teams{client: client}
}

// Retrieves list of team members for the given team
func (s *Teams) Members(teamName string) ([]*assist.User, error) {
	body, err := s.client.get("/teams/" + teamName + "/members")
	if err != nil {
		return nil, err
	}

	members := make([]*assist.User, 0)
	jsonErr := json.Unmarshal(body, &members)
	return members, jsonErr
}

// Retrieves list of team shots for the given team
func (s *Teams) Shots(teamName string) ([]*assist.Shot, error) {
	body, err := s.client.get("/teams/" + teamName + "/shots")
	if err != nil {
		return nil, err
	}

	shots := make([]*assist.Shot, 0)
	jsonErr := json.Unmarshal(body, &shots)
	return shots, jsonErr
}
