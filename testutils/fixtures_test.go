package testutils

import (
	"fmt"
	"strings"
	"testing"
)

func TestUrlToFixturePath(t *testing.T) {

	inputs := []struct {
		url  string
		path string
		root string
	}{
		{"/teams/simplebits/shots", "fixtures/teams-simplebits-shots.json", "./"},
	}

	for _, input := range inputs {
		path := UrlToFixturePath(input.root, input.url)
		if path != input.path {
			t.Errorf("%s != %s | url was: %s", path, input.path, input.url)
		}
	}

}
