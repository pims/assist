package assist

import (
	"encoding/json"
	"fmt"
)

type ProjectsService struct {
	client *Client
}

func NewProjectsService(client *Client) *ProjectsService {
	return &ProjectsService{client: client}
}

func (s *ProjectsService) Get(id int) (*Project, error) {
	body, err := s.client.get(fmt.Sprintf("/projects/%d", id))
	if err != nil {
		return nil, err
	}
	project := &Project{}
	jsonErr := json.Unmarshal(body, project)
	return project, jsonErr
}

func (s *ProjectsService) Shots(id int) ([]*Shot, error) {
	return s.client.shots(fmt.Sprintf("/projects/%d/shots", id))
}
