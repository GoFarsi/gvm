package gvm

import (
	"context"
	"fmt"
	"github.com/GoFarsi/gvm/cli"
	"github.com/GoFarsi/gvm/internal"
	"github.com/GoFarsi/gvm/internal/tags"
	"github.com/hashicorp/go-version"
	"strings"
	"time"
)

const (
	gvmTags     = "https://api.github.com/repos/GoFarsi/gvm/git/refs/tags"
	releasePath = "https://github.com/GoFarsi/gvm/releases/tag/v"
	refPrefix   = "refs/tags/v"
)

type gvm struct {
	tags []*tags.Tag
}

func NewGVMCheck() (*gvm, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	tags, err := cli.NewGetRequest[[]*tags.Tag](ctx, gvmTags, "")
	if err != nil {
		return nil, err
	}
	return &gvm{tags}, nil
}

func (g *gvm) CheckNewVersion() (string, error) {
	ver := &version.Version{}
	if len(g.tags) != 0 {
		if strings.HasPrefix(g.tags[len(g.tags)-1].Ref, refPrefix) {
			latestVer, err := version.NewVersion(g.tags[len(g.tags)-1].Ref[11:])
			if err != nil {
				return "", err
			}
			ver = latestVer
		}
	} else {
		return "", nil
	}

	gvmVer, err := internal.GvmVersion.Version()
	if err != nil {
		return "", err
	}

	if ver.GreaterThan(gvmVer) {
		return fmt.Sprintf("new version %s is available for download: %s%s", ver.String(), releasePath, ver.String()), nil
	}

	return "your gvm is latest version.", nil
}
