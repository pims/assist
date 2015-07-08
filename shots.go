package assist

type ShotsService struct {
	client *Client
}

type QueryParams struct {
	List      *string
	TimeFrame *string
	Date      *string
	Sort      *string
}

type CreateShot struct {
	Title           string
	Image           string // It must be exactly 400x300 or 800x600, no larger than eight megabytes, and be a GIF, JPG, or PNG.
	Description     string
	Tags            []string // Limited to a maximum of 12 tags.
	TeamId          *int
	ReboundSourceId *int
}

type UpdateShot struct {
	Title       string
	Description string
	Tags        []string // Limited to a maximum of 12 tags.
	TeamId      *int
}

func (s *ShotsService) List(params *QueryParams) ([]*Shot, error) {
	return s.client.shots("/shots")
}

func (s *ShotsService) Get(id int) (*Shot, error) {
	return s.client.shot("/shot/" + string(id))
}

func (s *ShotsService) Create(name, description string) (*Shot, error) {
	return nil, nil
}

func (s *ShotsService) Update(id int, name, description string) (*Shot, error) {
	return nil, nil
}

func (s *ShotsService) Delete(id int) error {
	return nil
}
