package updater

import (
	"fmt"
	"net/http"

	"github.com/antonholmquist/jason"
	"github.com/mcuadros/go-version"
	"github.com/pkg/errors"
)

// CheckWithTag using release tag
func CheckWithTag(organizer, project, now string) (string, error) {
	return Check(organizer, project, now, "tag_name")
}

// CheckWithTitle using release title
func CheckWithTitle(organizer, project, now string) (string, error) {
	return Check(organizer, project, now, "name")
}

// Check using key
func Check(organizer, project, now, key string) (string, error) {
	api := fmt.Sprintf("https://api.github.com/repos/%s/%s/releases/latest", organizer, project)
	res, err := http.Get(api)
	if err != nil {
		return "", errors.Wrap(err, "failed to get request: "+api)
	}
	defer res.Body.Close()

	json, err := jason.NewObjectFromReader(res.Body)
	if err != nil {
		return "", errors.Wrap(err, "failed to parse json")
	}

	isnew, err := json.GetString(key)
	if err != nil {
		return "", errors.Wrap(err, "failed to get value related to version from json")
	}

	if version.Compare(now, isnew, "<") {
		return fmt.Sprintf("update available. version: %s", isnew), nil
	}

	return fmt.Sprintf("update not available."), nil
}
