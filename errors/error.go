package errors

import "errors"

var (
	ERR_NETWORK_TIMEOUT = errors.New("cannot connect to network for get list of go versions")
	ERR_INVALID_VALUE   = errors.New("invalid value of switch")
)
