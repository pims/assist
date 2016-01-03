package assist

import (
	"fmt"
)

type BucketsService struct {
	client *Client
}

// NewBucketService creates a new BucketService client
func NewBucketsService(client *Client) *BucketsService {
	return &BucketsService{client: client}
}

func (s *BucketsService) Get(id int) (*Bucket, error) {
	return s.client.bucket(fmt.Sprintf("/buckets/%d", id))
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
