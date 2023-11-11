package hostsfile

import "net"

// IsCurrentHost returns true if the host is the current host.
func (h *Host) IsCurrentHost() bool {
	return h.address == "127.0.0.1" || h.address == "::1"
}

// IsPrivateHost returns true if the host is a private host.
// rfc: https://datatracker.ietf.org/doc/html/rfc1918
func (h *Host) IsPrivateHost() bool {
	ip := net.ParseIP(h.address)
	if ip == nil {
		return false
	}

	return ip.IsPrivate()
}

// IsPublicHost returns true if the host is a public host.
func (h *Host) IsPublicHost() bool {
	return !h.IsPrivateHost()
}

// HasHostName returns true if the host has the hostname.
func (h *Host) HasHostName(hostname string) bool {
	for _, d := range h.hostNames {
		if d == hostname {
			return true
		}
	}
	return false
}
