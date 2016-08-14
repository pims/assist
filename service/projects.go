package service

import (
	"encoding/json"
	"fmt"
	"github.com/pims/assist"
)

type Projects struct {
	client *Client
}

func NewProjects(client *Client) *Projects {
	return &Projects{client: client}
}

func (s *Projects) Get(id int) (*assist.Project, error) {
	body, err := s.client.get(fmt.Sprintf("/projects/%d", id))
	if err != nil {
		return nil, err
	}
	project := &assist.Project{}
	jsonErr := json.Unmarshal(body, project)
	return project, jsonErr
}

func (s *Projects) Shots(id int) ([]*assist.Shot, error) {
	return s.client.shots(fmt.Sprintf("/projects/%d/shots", id))
}
