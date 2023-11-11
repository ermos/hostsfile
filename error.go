package hostsfile

import (
	"errors"
)

var (
	ErrHostAddressIsRequired         = errors.New("host address is required")
	ErrHostRequireAtLeastOneHostname = errors.New("host require at least one hostname")
	ErrUnsupportedOperatingSystem    = errors.New("unsupported operating system")
	ErrHostNotFound                  = errors.New("host not found")
)
