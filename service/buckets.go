package service

import (
	"encoding/json"
	"fmt"
	"github.com/pims/assist"
)

// Buckets is something
type Buckets struct {
	client *Client
}

// NewBuckets creates a new Bucket service
func NewBuckets(client *Client) *Buckets {
	return &Buckets{client: client}
}

// Get a Bucket by id
func (s *Buckets) Get(id int) (*assist.Bucket, error) {
	body, err := s.client.get(fmt.Sprintf("/buckets/%d", id))
	if err != nil {
		return nil, err
	}
	bucket := &assist.Bucket{}
	jsonErr := json.Unmarshal(body, bucket)
	return bucket, jsonErr
}

func (s *Buckets) Create(name, description string) (*assist.Bucket, error) {
	return nil, assist.ErrNotImplemented
}

func (s *Buckets) Update(id int, name, description string) (*assist.Bucket, error) {
	return nil, assist.ErrNotImplemented
}

func (s *Buckets) Delete(id int) (*assist.Bucket, error) {
	return nil, assist.ErrNotImplemented
}

func (s *Buckets) Shots(id int) ([]*assist.Shot, error) {
	return s.client.shots(fmt.Sprintf("/buckets/%d/shots", id))
}

func (s *Buckets) Add(id string, shotId int) error {
	return assist.ErrNotImplemented
}
