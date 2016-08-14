package assist

import (
	"encoding/json"
)

// Links represent a map from name -> http urls to a user's social network links
type Links map[string]string

// User represents a dribbble user
type User struct {
	Id                    int    `json:"id"`
	Name                  string `json:"name"`
	Username              string `json:"username"`
	HtmlUrl               string `json:"html_url"`
	AvatarUrl             string `json:"avatar_url"`
	Bio                   string `json:"bio"`
	Location              string `json:"location"`
	Links                 Links  `json:"links"`
	BucketsCount          int    `json:"buckets_count"`
	CommentsReceivedCount int    `json:"comments_received_count"`
	FollowersCount        int    `json:"followers_count"`
	FollowingsCount       int    `json:"followings_count"`
	LikesCount            int    `json:"likes_count"`
	LikesReceivedCount    int    `json:"likes_received_count"`
	ProjectsCount         int    `json:"projects_count"`
	ReboundsReceivedCount int    `json:"rebounds_count_received"`
	ShotsCount            int    `json:"shots_count"`
	TeamsCount            int    `json:"teams_count"`
	CanUploadShot         bool   `json:"can_upload_shot"`
	Type                  string `json:"type"`
	Pro                   bool   `json:"pro"`
	BucketsUrl            string `json:"buckets_url"`
	FollowersUrl          string `json:"followers_url"`
	FollowingUrl          string `json:"following_url"`
	LikesUrl              string `json:"likes_url"`
	ShotsUrl              string `json:"shots_url"`
	TeamsUrl              string `json:"teams_url"`
	CreatedAt             string `json:"created_at"`
	UpdatedAt             string `json:"updated_at"`
}

func (u *User) String() string {
	buff, _ := json.Marshal(u)
	return string(buff)
}

// Team is a strict superset of a User
type Team struct {
	User

	MembersCount int    `json:"members_count"`
	MembersUrl   string `json:"members_url"`
	TeamShotsUrl string `json:"team_shots_url"`
}

func (t *Team) String() string {
	buff, _ := json.Marshal(t)
	return string(buff)
}

// Shot represents a dribbble shot
type Shot struct {
	Id               int               `json:"id"`
	Title            string            `json:"title"`
	Description      string            `json:"description"`
	Width            int               `json:"width"`
	Height           int               `json:"height"`
	Images           map[string]string `json:"images"`
	ViewsCount       int               `json:"views_count"`
	LikesCount       int               `json:"likes_count"`
	CommentsCount    int               `json:"comments_count"`
	AttachmentsCount int               `json:"attachments_count"`
	ReboundsCount    int               `json:"rebounds_count"`
	BucketsCount     int               `json:"buckets_count"`
	CreatedAt        string            `json:"created_at"`
	UpdatedAt        string            `json:"updated_at"`
	HtmlUrl          string            `json:"html_url"`
	AttachmentsUrl   string            `json:"attachments_url"`
	BucketsUrl       string            `json:"buckets_url"`
	CommentsUrl      string            `json:"comments_url"`
	LikesUrl         string            `json:"likes_url"`
	ProjectsUrl      string            `json:"projects_url"`
	ReboundsUrl      string            `json:"rebounds_url"`
	Tags             []string          `json:"tags"`
	Team             *Team             `json:"team"`
	User             *User             `json:"user"`
}

func (s *Shot) String() string {
	buff, _ := json.Marshal(s)
	return string(buff)
}

// Project represents a dribbble project
type Project struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	ShotsCount  int    `json:"shots_count"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
	User        User   `json:"user"`
}

func (p *Project) String() string {
	buff, _ := json.Marshal(p)
	return string(buff)
}

// Bucket represents a dribbble bucket
type Bucket struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	ShotsCount  int    `json:"shots_count"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
	User        User   `json:"user"`
}

func (b *Bucket) String() string {
	buff, _ := json.Marshal(b)
	return string(buff)
}

// Like represents a dribbble like
type Like struct {
	Id        int    `json:"id"`
	CreatedAt string `json:"created_at"`
	Shot      *Shot  `json:"shot"`
}

// Follower represents a dribbble follower
type Follower struct {
	Id        int    `json:"id"`
	CreatedAt string `json:"created_at"`
	User      *User  `json:"follower"`
}

// Following represents a dribbble followee
type Following struct {
	Id        int    `json:"id"`
	CreatedAt string `json:"created_at"`
	User      *User  `json:"followee"`
}

type ContentType string

type Attachment struct {
	Id           int         `json:"id"`
	Url          string      `json:"url"`
	ThumbnailUrl string      `json:"thumbnail_url"`
	Size         int         `json:"size"`
	ContentType  ContentType `json:"content_type"`
	ViewsCount   int         `json:"views_count"`
	CreatedAt    string      `json:"created_at"`
}

type Comment struct {
	Id         int    `json:"id"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
	Body       string `json:"body"`
	LikesCount int    `json:"likes_count"`
	LikesUrl   string `json:"likes_url"`
}
