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

func TestShotBuckets(t *testing.T) {
	ts, client := configureClient("testutils", t)
	defer ts.Close()

	shotId := 471756
	buckets, err := client.Shots.Buckets(shotId)
	if err != nil {
		t.Errorf("Failed: %v", err)
		t.FailNow()
	}

	if len(buckets) != 1 {
		t.Error("Teams length don't match")
		t.FailNow()
	}

	if buckets[0].Id != 2754 {
		t.Errorf("Got %d, expected: %d", buckets[0].Id, 2754)
		t.FailNow()
	}

	if buckets[0].User.Id != 1 {
		t.Error("User id for bucket doesn't match")
		t.FailNow()
	}
}

func TestBucket(t *testing.T) {
	ts, client := configureClient("testutils", t)
	defer ts.Close()

	bucketId := 2754
	bucket, err := client.Buckets.Get(bucketId)
	if err != nil {
		t.Errorf("Failed: %v", err)
		t.FailNow()
	}

	if bucket.Id != 2754 {
		t.Errorf("Got %d, expected: %d", bucket.Id, 2754)
		t.FailNow()
	}

	if bucket.User.Id != 1 {
		t.Error("User id for bucket doesn't match")
		t.FailNow()
	}
}

func TestBucketShots(t *testing.T) {
	ts, client := configureClient("testutils", t)
	defer ts.Close()

	shots, err := client.Buckets.Shots(2754)
	if err != nil {
		t.Errorf("Failed: %v", err)
		t.FailNow()
	}

	if len(shots) != 1 {
		t.Error("Length don't match")
		t.FailNow()
	}

	if shots[0].LikesCount != 149 {
		t.Errorf("Got %d, expected: %d", shots[0].LikesCount, 149)
		t.FailNow()
	}
	if shots[0].User.Id != 1 {
		t.Errorf("Got %d, expected: %d", shots[0].User.Id, 1)
		t.FailNow()
	}

	if shots[0].Team.Id != 39 {
		t.Errorf("Got %d, expected: %d", shots[0].Team.Id, 39)
		t.FailNow()
	}
}

func TestProject(t *testing.T) {
	ts, client := configureClient("testutils", t)
	defer ts.Close()

	projectId := 3
	project, err := client.Projects.Get(projectId)
	if err != nil {
		t.Errorf("Failed: %v", err)
		t.FailNow()
	}

	checkInt(t, project.Id, projectId)
	checkInt(t, project.ShotsCount, 4)
	checkInt(t, project.ShotsCount, 4)
	checkInt(t, project.User.ProjectsCount, 8)
}

func checkInt(t *testing.T, got, want int) {
	if got != want {
		t.Errorf("Got %d, expected: %d", got, want)
		t.FailNow()
	}
}
