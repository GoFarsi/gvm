package internal

import (
	"fmt"
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
