package service

import (
	"encoding/json"
	"github.com/pims/assist"
)

type Users struct {
	client *Client
}

func NewUsers(client *Client) *Users {
	return &Users{client: client}
}

func (s *Users) Get(username string) (*assist.User, error) {
	body, err := s.client.get("/users/" + username)
	if err != nil {
		return nil, err
	}
	user := &assist.User{}
	jsonErr := json.Unmarshal(body, user)
	return user, jsonErr
}

func (s *Users) Likes(name string) ([]*assist.Shot, error) {
	body, err := s.client.get("/users/" + name + "/likes")
	if err != nil {
		return nil, err
	}

	likes := make([]*assist.Like, 0)
	jsonErr := json.Unmarshal(body, &likes)

	shots := make([]*assist.Shot, len(likes))
	for i, like := range likes {
		shots[i] = like.Shot
	}
	return shots, jsonErr
}

func (s *Users) Buckets(name string) ([]*assist.Bucket, error) {
	body, err := s.client.get("/users/" + name + "/buckets")
	if err != nil {
		return nil, err
	}

	buckets := make([]*assist.Bucket, 0)
	jsonErr := json.Unmarshal(body, &buckets)
	return buckets, jsonErr
}

func (s *Users) Followers(name string) ([]*assist.User, error) {
	body, err := s.client.get("/users/" + name + "/followers")
	if err != nil {
		return nil, err
	}

	followers := make([]*assist.Follower, 0)
	jsonErr := json.Unmarshal(body, &followers)

	users := make([]*assist.User, len(followers))
	for i, follower := range followers {
		users[i] = follower.User
	}
	return users, jsonErr
}

func (s *Users) Following(name string) ([]*assist.User, error) {
	body, err := s.client.get("/users/" + name + "/following")
	if err != nil {
		return nil, err
	}

	followings := make([]*assist.Following, 0)
	jsonErr := json.Unmarshal(body, &followings)
	users := make([]*assist.User, len(followings))
	for i, following := range followings {
		users[i] = following.User
	}
	return users, jsonErr
}

func (s *Users) Follow(name string) error {
	return assist.ErrNotImplemented
}

func (s *Users) Unfollow(name string) error {
	return s.client.delete("/users/" + name + "/follow")
}

func (s *Users) Friend(name string) (bool, error) {
	resp, err := s.client.rawGet("/user/following/" + name)
	return resp.StatusCode == 204, err
}

func (s *Users) Friends(name, target string) (bool, error) {
	resp, err := s.client.rawGet("/users/" + name + "/following/" + target)
	return resp.StatusCode == 204, err
}

func (s *Users) Me() (*assist.User, error) {
	body, err := s.client.get("/user")
	if err != nil {
		return nil, err
	}
	user := &assist.User{}
	jsonErr := json.Unmarshal(body, user)
	return user, jsonErr
}

func (s *Users) Projects(name string) ([]*assist.Project, error) {
	body, err := s.client.get("/users/" + name + "/projects")
	if err != nil {
		return nil, err
	}
	projects := make([]*assist.Project, 0)
	jsonErr := json.Unmarshal(body, projects)
	return projects, jsonErr
}

func (s *Users) Teams(name string) ([]*assist.Team, error) {
	body, err := s.client.get("/users/" + name + "/teams")
	if err != nil {
		return nil, err
	}
	teams := make([]*assist.Team, 0)
	jsonErr := json.Unmarshal(body, &teams)
	return teams, jsonErr
}
