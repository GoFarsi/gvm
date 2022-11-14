package internal

import (
	"fmt"
	vers "github.com/hashicorp/go-version"
	"github.com/martinlindhe/notify"
	"strings"
)

func GetDocWithSpecificVersion(ver string) string {
	return releaseDoc + ver
}

func GetFullDoc() (string, error) {
	list, err := NewList()
	if err != nil {
		return "", err
	}

	vers := list.GetVersions()

	var b strings.Builder
	for _, v := range vers {
		if len(v) > 4 {
			continue
		}
		fmt.Fprintf(&b, "go%s\t\t%s\n", v, releaseDoc+v)
	}
	return b.String(), nil
}

func ReleaseNotification() error {
	list, err := NewList()
	if err != nil {
		return err
	}

	installedVersion, ok := checkGoInstalled()
	if !ok {
		return nil
	}

	remoteVersion, err := vers.NewVersion(list.LastVersion())
	if err != nil {
		return err
	}

	localVersion, err := vers.NewVersion(installedVersion)
	if err != nil {
		return err
	}

	if remoteVersion.GreaterThan(localVersion) {
		notify.Notify("gvm", fmt.Sprintf("gvm: Go %s has been released", remoteVersion.String()), fmt.Sprintf("A new version of Go %s has been released, and you can upgrade to latest version using the 'gvm upgrade' command.", remoteVersion.String()), "")
	}

	return nil
}
