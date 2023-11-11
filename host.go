package hostsfile

import (
	"fmt"
	"strings"
)

// Host represents a host in the hosts file.
type Host struct {
	address   string
	hostNames []string
	comment   string
	parent    *Hosts
}

// GetAddress returns the address of the host.
func (h *Host) GetAddress() string {
	return h.address
}

// GetHostNames returns the hostname list of the host.
func (h *Host) GetHostNames() []string {
	return h.hostNames
}

// GetComment returns the comment of the host.
func (h *Host) GetComment() string {
	return h.comment
}

// ToString returns the host as a string.
func (h *Host) ToString() (string, error) {
	var lineBuilder []string

	if h.address == "" {
		return "", ErrHostAddressIsRequired
	}

	lineBuilder = append(lineBuilder, h.address)

	if len(h.hostNames) == 0 {
		return "", ErrHostRequireAtLeastOneHostname
	}

	lineBuilder = append(lineBuilder, h.hostNames...)

	if h.comment != "" {
		lineBuilder = append(lineBuilder, fmt.Sprintf("# %s", h.comment))
	}

	return strings.Join(lineBuilder, " "), nil
}
