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

// Default API endpoint for Dribbble.com
const DefaultAPIEndpoint = "https://api.dribbble.com/v1"

var (
	// ErrNotImplemented for methods not yet implemented
	ErrNotImplemented = errors.New("method not implemented.")
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

// DribbbleError Error returned by Dribbble API
type DribbbleError struct {
	Attribute string `json:"attribute"`
	Message   string `json:"message"`
}

func (d *DribbbleError) String() string {
	return fmt.Sprintf("%s %s", d.Message, d.Attribute)
}

// Client contains methods to query
// the Dribbble API
type Client struct {
	client             *http.Client
	config             *Config
	RateLimitRemaining int
	Shots              *ShotsService
	Users              *UsersService
	Teams              *TeamsService
	Buckets            *BucketsService
	Projects           *ProjectsService
	Comments           *CommentsService
	logger             *log.Logger
	sync.Mutex
}

// Config for Client
type Config struct {
	Token    string
	Endpoint string
	logger   *log.Logger
}

var defaultLogger = log.New(os.Stderr, "[assist] ", log.LstdFlags)

// NewConfig creates new Config
func NewConfig(token, endpoint string) *Config {
	return &Config{Token: token, Endpoint: endpoint, logger: defaultLogger}
}

// NewClient creates new client with given configuration
func NewClient(config *Config) *Client {
	c := &Client{
		client: http.DefaultClient,
		config: config,
	}

	return configure(c, config.logger)
}

func configure(client *Client, logger *log.Logger) *Client {
	client.logger = logger
	client.Shots = NewShotsService(client)
	client.Users = NewUsersService(client)
	client.Teams = NewTeamsService(client)
	client.Buckets = NewBucketsService(client)
	client.Projects = NewProjectsService(client)
	client.Comments = NewCommentsService(client)
	return client
}

// NewDefaultClient creates new client with default configuration
func NewDefaultClient() *Client {
	c := &Client{
		client: http.DefaultClient,
		config: &Config{
			Token:    os.Getenv("DRIBBBLE_TOKEN"),
			Endpoint: DefaultAPIEndpoint,
		},
	}
	return configure(c, defaultLogger)
}

// Status retrieves most recent rate limit remaining
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
	if err != nil {
		return nil, err
	}

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
		c.logger.Println(err)
		return []byte{}, err
	}

	defer resp.Body.Close()
	bodyContent, err := ioutil.ReadAll(resp.Body)

	if c := resp.StatusCode; 200 <= c && c <= 299 {
		return bodyContent, nil
	}

	httpErr := &DribbbleError{}
	jsonErr := json.Unmarshal(bodyContent, httpErr)

	if jsonErr != nil {
		c.logger.Println(jsonErr)
		return []byte{}, jsonErr
	}
	c.logger.Println(httpErr)
	return []byte{}, errors.New(httpErr.Message)

}

// Raw HTTP GET wrapper
func (c *Client) rawGet(path string) (*http.Response, error) {
	req, _ := http.NewRequest("GET", c.url(path), nil)
	return c.do(req)
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

// Convenience method to decode response as []*Bucket
func (c *Client) buckets(path string) ([]*Bucket, error) {
	body, err := c.get(path)
	if err != nil {
		return nil, err
	}
	collection := make([]*Bucket, 0)
	jsonErr := json.Unmarshal(body, &collection)
	return collection, jsonErr
}

// Convenience method to decode response as *Bucket
func (c *Client) bucket(path string) (*Bucket, error) {
	body, err := c.get(path)
	if err != nil {
		return nil, err
	}
	bucket := &Bucket{}
	jsonErr := json.Unmarshal(body, bucket)
	return bucket, jsonErr
}
