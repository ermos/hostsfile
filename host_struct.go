package hostsfile

import "net"

// Host represents a host in the hosts file.
type Host struct {
	// Address is the IP address of the host.
	Address string
	// HostNames is the hostnames of the host.
	HostNames []string
	// Comment is the comment of the host.
	Comment string
}

// IsCurrentHost returns true if the host is the current host.
func (h Host) IsCurrentHost() bool {
	return h.Address == "127.0.0.1" || h.Address == "::1"
}

// IsPrivateHost returns true if the host is a private host.
// rfc: https://datatracker.ietf.org/doc/html/rfc1918
func (h Host) IsPrivateHost() bool {
	ip := net.ParseIP(h.Address)
	if ip == nil {
		return false
	}

	return ip.IsPrivate()
}

// IsPublicHost returns true if the host is a public host.
func (h Host) IsPublicHost() bool {
	return !h.IsPrivateHost()
}

// HasHostName returns true if the host has the hostname.
func (h Host) HasHostName(hostname string) bool {
	for _, d := range h.HostNames {
		if d == hostname {
			return true
		}
	}
	return false
}
