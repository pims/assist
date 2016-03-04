package assist

import (
	"encoding/json"
)

// BucketService is something
type BucketService struct {
	client *Client
}

// Get a Bucket by id
func (s *BucketService) Get(id int) (*Bucket, error) {
	body, err := s.client.get("/buckets/" + string(id))
	if err != nil {
		return nil, err
	}
	bucket := &Bucket{}
	jsonErr := json.Unmarshal(body, bucket)
	return bucket, jsonErr
}

func (s *BucketService) Create(name, description string) (*Bucket, error) {
	return nil, ErrNotImplemented
}

func (s *BucketService) Update(id int, name, description string) (*Bucket, error) {
	return nil, ErrNotImplemented
}

func (s *BucketService) Delete(id int) (*Bucket, error) {
	return nil, ErrNotImplemented
}

func (s *BucketService) Shots(id int) ([]*Shot, error) {
	return s.client.shots("/buckets/" + string(id) + "/shots")
}

func (s *BucketService) Add(id string, shotId int) error {
	return ErrNotImplemented
}
