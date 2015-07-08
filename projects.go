package assist

import (
	"encoding/json"
)

type ProjectsService struct {
	client *Client
}

func (s *ProjectsService) Get(id int) (*Project, error) {
	body, err := s.client.get("/projects/" + string(id))
	if err != nil {
		return nil, err
	}
	project := &Project{}
	jsonErr := json.Unmarshal(body, project)
	return project, jsonErr
}

func (s *ProjectsService) Shots(id int) ([]*Shot, error) {
	return s.client.shots("/projects/" + string(id) + "/shots")
}
