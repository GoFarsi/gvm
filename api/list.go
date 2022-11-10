package api

import (
	"context"
	"fmt"
	"github.com/GoFarsi/gvm/cli"
	"reflect"
	"strings"
	"time"
)

type List struct {
	tags     []*tag
	versions []*Version
}

type Version struct {
	Version string
	Commit  string
}

type tag struct {
	Ref    string `json:"ref"`
	NodeId string `json:"node_id"`
	Url    string `json:"url"`
	Object *struct {
		Sha  string `json:"sha"`
		Type string `json:"type"`
		Url  string `json:"url"`
	} `json:"object"`
}

func NewList() (*List, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	tags, err := cli.NewGetRequest[[]*tag](ctx, tagsAPI, "")
	if err != nil {
		return nil, err
	}
	return &List{
		tags:     tags,
		versions: make([]*Version, 0),
	}, nil
}

func (t *List) Print() string {
	if len(t.versions) == 0 {
		t.getVersions()
	}
	t.swap()
	return fmt.Sprintf("%s%v", t.header(), t.list())
}

func (t *List) GetVersions() []string {
	vers := []string{}
	if len(t.versions) == 0 {
		t.getVersions()
	}
	t.swap()

	for _, version := range t.versions {
		vers = append(vers, version.Version)
	}
	return vers
}

func (t *List) LastVersion() string {
	return t.GetVersions()[0]
}

func (t *List) SetNumOfVersions(numOfVersion int) {
	if len(t.versions) == 0 {
		t.getVersions()
	}
	t.versions = t.versions[len(t.versions)-numOfVersion:]
}

func (t *List) list() string {
	var b strings.Builder
	for _, ver := range t.versions {
		fmt.Fprintf(&b, "%s\t\t%s\n", ver.Version, ver.Commit)
	}
	return b.String()
}

func (t *List) swap() {
	swp := reflect.Swapper(t.versions)

	for i := 0; i < len(t.versions)/2; i++ {
		swp(i, len(t.versions)-1-i)
	}
}

func (t *List) getVersions() {
	for _, tag := range t.tags {
		ver := &Version{}
		if strings.HasPrefix(tag.Ref, refPrefix) {
			ver.Version = tag.Ref[12:]
			if tag.Object != nil {
				ver.Commit = tag.Object.Sha
			}
			t.versions = append(t.versions, ver)
		}

	}
}

func (t *List) header() string {
	return fmt.Sprintf("Version\t\tCommit\n=======\t\t========================================\n")
}
