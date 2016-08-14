package assist

import (
	"encoding/json"
	"fmt"
)

// BucketsService is something
type BucketsService struct {
	client *Client
}

// NewProjectsService creates a new
func NewBucketsService(client *Client) *BucketsService {
	return &BucketsService{client: client}
}

// Get a Bucket by id
func (s *BucketsService) Get(id int) (*Bucket, error) {
	body, err := s.client.get(fmt.Sprintf("/buckets/%d", id))
	if err != nil {
		return nil, err
	}
	bucket := &Bucket{}
	jsonErr := json.Unmarshal(body, bucket)
	return bucket, jsonErr
}

func (s *BucketsService) Create(name, description string) (*Bucket, error) {
	return nil, ErrNotImplemented
}

func (s *BucketsService) Update(id int, name, description string) (*Bucket, error) {
	return nil, ErrNotImplemented
}

func (s *BucketsService) Delete(id int) (*Bucket, error) {
	return nil, ErrNotImplemented
}

func (s *BucketsService) Shots(id int) ([]*Shot, error) {
	return s.client.shots(fmt.Sprintf("/buckets/%d/shots", id))
}

func (s *BucketsService) Add(id string, shotId int) error {
	return ErrNotImplemented
}
