package tests

import (
	"fmt"
	"strings"
	"testing"
)

func urlToFixturePath(url string) string {
	stuff := strings.TrimPrefix(url, "/")
	path := strings.Replace(stuff, "/", "-", -1)
	return fmt.Sprintf("fixtures/%s.json", path)
}

func TestUrlToFixturePath(t *testing.T) {

	inputs := []struct {
		url  string
		path string
	}{
		{"/teams/simplebits/shots", "fixtures/teams-simplebits-shots.json"},
	}

	for _, input := range inputs {
		path := urlToFixturePath(input.url)
		if path != input.path {
			t.Errorf("%s != %s | url was: %s", path, input.path, input.url)
		}
	}

}
