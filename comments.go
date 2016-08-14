package assist

import (
	"encoding/json"
	"fmt"
)

type CommentsService struct {
	client *Client
}

func NewCommentsService(client *Client) *CommentsService {
	return &CommentsService{client: client}
}

func (s *CommentsService) Get(shotId, commentId int) (*Comment, error) {
	body, err := s.client.get(fmt.Sprintf("/shots/%d/comments/%d", shotId, commentId))
	if err != nil {
		return nil, err
	}
	comment := &Comment{}
	jsonErr := json.Unmarshal(body, comment)
	return comment, jsonErr
}

func (s *CommentsService) Shot(id int) ([]*Comment, error) {
	body, err := s.client.get(fmt.Sprintf("/shots/%d/comments", id))
	if err != nil {
		return nil, err
	}
	comments := make([]*Comment, 0)
	jsonErr := json.Unmarshal(body, &comments)
	return comments, jsonErr
}

func (s *CommentsService) Likes(shotId, commentId int) ([]*Like, error) {
	body, err := s.client.get(fmt.Sprintf("/shots/%d/comments/%d/likes", shotId, commentId))
	if err != nil {
		return nil, err
	}
	likes := make([]*Like, 0)
	jsonErr := json.Unmarshal(body, &likes)
	return likes, jsonErr
}
