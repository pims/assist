package assist

type UserService interface {
	Get(username string) (*User, error)
	Likes(name string) ([]*Shot, error)
	Buckets(name string) ([]*Bucket, error)
	Followers(name string) ([]*User, error)
	Following(name string) ([]*User, error)
	Follow(name string) error
	Unfollow(name string) error
	Friend(name string) (bool, error)
	Friends(name, target string) (bool, error)
	Me() (*User, error)
	Projects(name string) ([]*Project, error)
	Teams(name string) ([]*Team, error)
}

type CommentService interface {
	Get(shotId, commentId int) (*Comment, error)
	Shot(id int) ([]*Comment, error)
	Likes(shotId, commentId int) ([]*Like, error)
}

type ProjectService interface {
	Get(id int) (*Project, error)
	Shots(id int) ([]*Shot, error)
}

type ShotService interface {
	List(params *QueryParams) ([]*Shot, error)
	Get(id int) (*Shot, error)
	Create(name, description string) (*Shot, error)
	Update(id int, name, description string) (*Shot, error)
	Delete(id int) error
	Buckets(id int) ([]*Bucket, error)
	Attachments(id int) ([]*Attachment, error)
	Attachment(shotId, attachmentId int) (*Attachment, error)
}

type TeamService interface {
	Members(teamName string) ([]*User, error)
	// Retrieves list of team shots for the given team
	Shots(teamName string) ([]*Shot, error)
}

type BucketService interface {
	// Get a Bucket by id
	Get(id int) (*Bucket, error)
	Create(name, description string) (*Bucket, error)
	Update(id int, name, description string) (*Bucket, error)
	Delete(id int) (*Bucket, error)
	Shots(id int) ([]*Shot, error)
	Add(id string, shotId int) error
}
