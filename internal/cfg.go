package internal

import (
	"time"
)

const (
	prefix        = "go"
	refPrefix     = "refs/tags/go"
	filePattern   = "go%s.linux-%s.tar.gz"
	goInstallPath = "/usr/local/go"
	env           = "export PATH=$PATH:/usr/local/go/bin"
)

const (
	tagsAPI    = "https://api.github.com/repos/golang/go/git/refs/tags"
	releaseDoc = "https://go.dev/doc/go"
)

const (
	_defaultMirrorTimeout = 15 * time.Second
)

const (
	INSTALL   = "install"
	UPGRADE   = "upgrade"
	DOWNGRADE = "downgrade"
)

var mirrors = map[string]string{
	"default":   "https://go.dev/dl/",
	"gomirrors": "https://gomirrors.org/dl/go/",
}

var (
	warningInstalledGo = "Go installed with version %s in path %s, are you sure continue %s operate?"
	warningSudoAccess  = "For %s operation you need have 'sudo' access, are you sure have 'sudo' access?"
)
