package cli

import (
	"net/http"
)

func DefaultClient() *http.Client {
	return http.DefaultClient
}
