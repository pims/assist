package testutils

import (
	"fmt"
	"strings"
)

func UrlToFixturePath(rootPath, url string) string {
	stuff := strings.TrimPrefix(url, "/")
	path := strings.Replace(stuff, "/", "-", -1)
	return fmt.Sprintf(rootPath+"/fixtures/%s.json", path)
}
