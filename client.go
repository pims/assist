package assist

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"sync"
)

const DefaultApiEndpoint = "https://api.dribbble.com/v1"

var (
	// For methods not yet implemented
	ErrNotImplemented = errors.New("dribbble: method not implemented")
)

/*

curl -i https://api.dribbble.com/v1/users/simplebits

HTTP/1.1 200 OK
Date: Thu, 13 Feb 2014 19:30:30 GMT
ETag: "def2bc69c674e5b48cd281aa12c2c8e9"
Server: nginx
Status: 200 OK
Content-Type: application/json; charset=utf-8
Cache-Control: max-age=0, private, must-revalidate
X-RateLimit-Limit: 60
X-RateLimit-Remaining: 59
X-RateLimit-Reset: 1392321600

*/

// Error returned by Dribbble API
type DribbbleError struct {
	Attribute string `json:"attribute"`
	Message   string `json:"message"`
}

// Assist client contains methods to query
// the Dribbble API
type Client struct {
	client             *http.Client
	config             *Config
	RateLimitRemaining int
	Shots              *ShotsService
	Users              *UsersService
	sync.Mutex
}

type Config struct {
	Token    string
	Endpoint string
}

// Creates new Config
func NewConfig(token, endpoint string) *Config {
	return &Config{Token: token, Endpoint: endpoint}
}

// Creates new client with given configuration
func NewClient(config *Config) *Client {
	c := &Client{
		client: http.DefaultClient,
		config: config,
	}

	c.Shots = &ShotsService{client: c}
	c.Users = &UsersService{client: c}
	return c
}

// Creates new client with default configuration
func NewDefaultClient() *Client {
	c := &Client{
		client: http.DefaultClient,
		config: &Config{
			Token:    os.Getenv("DRIBBBLE_TOKEN"),
			Endpoint: DefaultApiEndpoint,
		},
	}

	c.Shots = &ShotsService{client: c}
	c.Users = &UsersService{client: c}
	return c
}

// Retrieve most recent rate limit remaining
func (c *Client) Status() int {
	c.Lock()
	defer c.Unlock()
	return c.RateLimitRemaining
}

// Convenience method to prepend configured endpoint to path
func (c *Client) url(path string) string {
	return c.config.Endpoint + path
}

// perform http request and keep track of ratelimit
func (c *Client) do(req *http.Request) (*http.Response, error) {
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.config.Token))
	resp, err := c.client.Do(req)
	val := resp.Header.Get("X-RateLimit-Remaining")
	if val != "" {
		c.Lock()
		i, _ := strconv.Atoi(val)
		c.RateLimitRemaining = i
		c.Unlock()
	}
	return resp, err
}

// Convenience HTTP GET wrapper
func (c *Client) get(path string) ([]byte, error) {
	req, _ := http.NewRequest("GET", c.url(path), nil)
	resp, err := c.do(req)
	if err != nil {
		log.Println(err)
		return []byte{}, err
	}

	defer resp.Body.Close()
	bodyContent, err := ioutil.ReadAll(resp.Body)

	// TODO: support additional valid responses HTTP 202, HTTP 204
	if resp.StatusCode != 200 {
		httpErr := &DribbbleError{}
		_ = json.Unmarshal(bodyContent, httpErr)
		return []byte{}, errors.New(httpErr.Message)
	}

	return bodyContent, nil
}

// Convenience HTTP PUT wrapper
func (c *Client) put(path string) error {
	return ErrNotImplemented
}

// Convenience HTTP POST wrapper
func (c *Client) post(path string) error {
	return ErrNotImplemented
}

// Convenience HTTP DELETE wrapper
func (c *Client) delete(path string) error {
	req, _ := http.NewRequest("DELETE", c.url(path), nil)
	_, err := c.do(req)
	return err
}

// Convenience method to decode response as []*Shot
func (c *Client) shots(path string) ([]*Shot, error) {
	body, err := c.get(path)
	if err != nil {
		return nil, err
	}
	collection := make([]*Shot, 0)
	jsonErr := json.Unmarshal(body, &collection)
	return collection, jsonErr
}

// Convenience method to decode response as *Shot
func (c *Client) shot(path string) (*Shot, error) {
	body, err := c.get(path)
	if err != nil {
		return nil, err
	}
	shot := &Shot{}
	jsonErr := json.Unmarshal(body, shot)
	return shot, jsonErr
}
