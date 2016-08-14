package service

import (
	"encoding/json"
	"fmt"
	"github.com/pims/assist"
)

type Comments struct {
	client *Client
}

func NewComments(client *Client) *Comments {
	return &Comments{client: client}
}

func (s *Comments) Get(shotId, commentId int) (*assist.Comment, error) {
	body, err := s.client.get(fmt.Sprintf("/shots/%d/comments/%d", shotId, commentId))
	if err != nil {
		return nil, err
	}
	comment := &assist.Comment{}
	jsonErr := json.Unmarshal(body, comment)
	return comment, jsonErr
}

func (s *Comments) Shot(id int) ([]*assist.Comment, error) {
	body, err := s.client.get(fmt.Sprintf("/shots/%d/comments", id))
	if err != nil {
		return nil, err
	}
	comments := make([]*assist.Comment, 0)
	jsonErr := json.Unmarshal(body, &comments)
	return comments, jsonErr
}

func (s *Comments) Likes(shotId, commentId int) ([]*assist.Like, error) {
	body, err := s.client.get(fmt.Sprintf("/shots/%d/comments/%d/likes", shotId, commentId))
	if err != nil {
		return nil, err
	}
	likes := make([]*assist.Like, 0)
	jsonErr := json.Unmarshal(body, &likes)
	return likes, jsonErr
}
