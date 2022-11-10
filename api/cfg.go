package api

import (
	"time"
)

const (
	prefix      = "go"
	refPrefix   = "refs/tags/go"
	filePattern = "go%s.linux-%s.tar.gz"
)

const (
	tagsAPI = "https://api.github.com/repos/golang/go/git/refs/tags"
)

const (
	_defaultMirrorTimeout = 15 * time.Second
)

var mirrors = map[string]string{
	"default":   "https://go.dev/dl/",
	"gomirrors": "https://gomirrors.org/dl/go/",
}
