package assist

import (
	"fmt"
	"github.com/pims/assist/testutils"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func configureClient(rootPath string, t *testing.T) (*httptest.Server, *Client) {

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fixturePath := testutils.UrlToFixturePath(rootPath, r.URL.String())
		buf, err := ioutil.ReadFile(fixturePath)
		if err != nil {
			t.Errorf("%v", err)
			t.FailNow()
		}
		r.Header.Add("X-RateLimit-Remaining", "10")
		fmt.Fprintf(w, "%s", buf)
	}))
	config := NewConfig("XXXX", ts.URL)
	return ts, NewClient(config)
}

func TestTeamsShots(t *testing.T) {
	ts, client := configureClient("testutils", t)
	defer ts.Close()

	shots, err := client.Teams.Shots("simplebits")
	if err != nil {
		t.Errorf("Failed: %v", err)
		t.FailNow()
	}

	if len(shots) != 1 {
		t.Error("Length don't match")
		t.FailNow()
	}
}

func TestUserGet(t *testing.T) {
	ts, client := configureClient("testutils", t)
	defer ts.Close()

	username := "simplebits"
	user, err := client.Users.Get(username)
	if err != nil {
		t.Errorf("Failed: %v", err)
		t.FailNow()
	}

	if user.Username != username {
		t.Error("Usernames don't match")
		t.FailNow()
	}
}

func TestUserTeams(t *testing.T) {
	ts, client := configureClient("testutils", t)
	defer ts.Close()

	username := "simplebits"
	teams, err := client.Users.Teams(username)
	if err != nil {
		t.Errorf("Failed: %v", err)
		t.FailNow()
	}

	if len(teams) != 1 {
		t.Error("Teams length don't match")
		t.FailNow()
	}

	if teams[0].Username != "dribbble" {
		t.Errorf("Got %s, expected: %s", teams[0].Username, "dribbble")
		t.FailNow()
	}

	if len(teams[0].Links) != 2 {
		t.Error("Teams length don't match")
		t.FailNow()
	}
}
