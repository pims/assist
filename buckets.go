package assist

type BucketService struct {
	client *Client
}

func (s *BucketService) Get(id int) (*Bucket, error) {
	return s.client.bucket("/buckets/" + string(id))
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
