package internal

import "github.com/hashicorp/go-version"

type ver string

const GvmVersion ver = "1.2.1"

func (ver) Version() (*version.Version, error) {
	return version.NewVersion(GvmVersion.String())
}

func (ver) String() string {
	return string(GvmVersion)
}
