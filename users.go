package assist

import (
	"encoding/json"
)

type UsersService struct {
	client *Client
}

func (s *UsersService) Get(username string) (*User, error) {
	body, err := s.client.get("/users/" + username)
	if err != nil {
		return nil, err
	}
	user := &User{}
	jsonErr := json.Unmarshal(body, user)
	return user, jsonErr
}

func (s *UsersService) Likes(name string) ([]*Shot, error) {
	body, err := s.client.get("/users/" + name + "/likes")
	if err != nil {
		return nil, err
	}

	likes := make([]*Like, 0)
	jsonErr := json.Unmarshal(body, &likes)

	shots := make([]*Shot, len(likes))
	for i, like := range likes {
		shots[i] = like.Shot
	}
	return shots, jsonErr
}

func (s *UsersService) Buckets(name string) ([]*Bucket, error) {
	body, err := s.client.get("/users/" + name + "/buckets")
	if err != nil {
		return nil, err
	}

	buckets := make([]*Bucket, 0)
	jsonErr := json.Unmarshal(body, &buckets)
	return buckets, jsonErr
}

func (s *UsersService) Followers(name string) ([]*User, error) {
	body, err := s.client.get("/users/" + name + "/followers")
	if err != nil {
		return nil, err
	}

	followers := make([]*Follower, 0)
	jsonErr := json.Unmarshal(body, &followers)

	users := make([]*User, len(followers))
	for i, follower := range followers {
		users[i] = follower.User
	}
	return users, jsonErr
}

func (s *UsersService) Following(name string) ([]*User, error) {
	body, err := s.client.get("/users/" + name + "/following")
	if err != nil {
		return nil, err
	}

	followings := make([]*Following, 0)
	jsonErr := json.Unmarshal(body, &followings)
	users := make([]*User, len(followings))
	for i, following := range followings {
		users[i] = following.User
	}
	return users, jsonErr
}

func (s *UsersService) Follow(name string) error {
	return ErrNotImplemented
}

func (s *UsersService) Unfollow(name string) error {
	return s.client.delete("/users/" + name + "/follow")
}

func (s *UsersService) Friend(name string) (bool, error) {
	resp, err := s.client.rawGet("/user/following/" + name)
	return resp.StatusCode == 204, err
}

func (s *UsersService) Friends(name, target string) (bool, error) {
	resp, err := s.client.rawGet("/users/" + name + "/following/" + target)
	return resp.StatusCode == 204, err
}

func (s *UsersService) Me() (*User, error) {
	body, err := s.client.get("/user")
	if err != nil {
		return nil, err
	}
	user := &User{}
	jsonErr := json.Unmarshal(body, user)
	return user, jsonErr
}

func (s *UsersService) Projects(name string) ([]*Project, error) {
	body, err := s.client.get("/users/" + name + "/projects")
	if err != nil {
		return nil, err
	}
	projects := make([]*Project, 0)
	jsonErr := json.Unmarshal(body, projects)
	return projects, jsonErr
}

func (s *UsersService) Teams(name string) ([]*Team, error) {
	body, err := s.client.get("/users/" + name + "/teams")
	if err != nil {
		return nil, err
	}
	teams := make([]*Team, 0)
	jsonErr := json.Unmarshal(body, &teams)
	return teams, jsonErr
}
