package service

import (
	"encoding/json"
	"fmt"
	"github.com/pims/assist"
)

type Shots struct {
	client *Client
}

func NewShots(client *Client) *Shots {
	return &Shots{client: client}
}

func (s *Shots) List(params *assist.QueryParams) ([]*assist.Shot, error) {
	return s.client.shots("/shots")
}

func (s *Shots) Get(id int) (*assist.Shot, error) {
	return s.client.shot(fmt.Sprintf("/shot/%d", id))
}

func (s *Shots) Create(name, description string) (*assist.Shot, error) {
	return nil, assist.ErrNotImplemented
}

func (s *Shots) Update(id int, name, description string) (*assist.Shot, error) {
	return nil, assist.ErrNotImplemented
}

func (s *Shots) Delete(id int) error {
	return assist.ErrNotImplemented
}

func (s *Shots) Buckets(id int) ([]*assist.Bucket, error) {
	return s.client.buckets(fmt.Sprintf("/shots/%d/buckets", id))
}

func (s *Shots) Attachments(id int) ([]*assist.Attachment, error) {
	body, err := s.client.get(fmt.Sprintf("/shots/%d/attachments", id))
	if err != nil {
		return nil, err
	}

	attachments := make([]*assist.Attachment, 0)
	jsonErr := json.Unmarshal(body, &attachments)
	return attachments, jsonErr
}

func (s *Shots) Attachment(shotId, attachmentId int) (*assist.Attachment, error) {
	body, err := s.client.get(fmt.Sprintf("/shots/%d/attachments/%d", shotId, attachmentId))
	if err != nil {
		return nil, err
	}

	attachment := &assist.Attachment{}
	jsonErr := json.Unmarshal(body, attachment)
	return attachment, jsonErr
}
