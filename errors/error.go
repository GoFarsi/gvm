package errors

import "errors"

var (
	ERR_NETWORK_TIMEOUT         = errors.New("cannot connect to network for get list of go versions")
	ERR_INVALID_VALUE           = errors.New("invalid value of switch")
	ERR_CANT_FIND_ACTIVE_MIRROR = errors.New("can't find active mirror on your network")
	ERR_SUDO_ACCESS             = errors.New("for install,upgrade and downgrade please run gvm with sudo access")
	ERR_CANT_REMOVE_OLD_VERSION = errors.New("can't remove old version of go in /usr/local/go")
	ERR_CANT_SET_ENV            = errors.New("can't set environment variable in /home/{user}/.profile file")
	ERR_UPGRADE_VERSION         = errors.New("version request for upgrade should bigger than installed version of go")
	ERR_DOWNGRADE_VERSION       = errors.New("version request for downgrade should smaller than installed version of go")
	ERR_READ_VERSION_CODE       = errors.New("can't read version code")
)
