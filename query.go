package assist

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
