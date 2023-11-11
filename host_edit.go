package hostsfile

// AddHostName adds the hostname to the host.
func (h *Host) AddHostName(hostName string) {
	for _, d := range h.hostNames {
		if d == hostName {
			return
		}
	}

	h.hostNames = append(h.hostNames, hostName)
}

// RemoveHostName removes the hostname from the host.
func (h *Host) RemoveHostName(hostName string) {
	for i, d := range h.hostNames {
		if d == hostName {
			h.hostNames = append(h.hostNames[:i], h.hostNames[i+1:]...)
		}
	}
}

// SetComment sets the comment for the host.
func (h *Host) SetComment(comment string) {
	h.comment = comment
}

// SetAddress sets the address for the host.
func (h *Host) SetAddress(address string) {
	h.address = address
}

// Remove removes the host from the hosts file.
func (h *Host) Remove() {
	for i, host := range h.parent.hosts {
		if host == h {
			h.parent.hosts = append(h.parent.hosts[:i], h.parent.hosts[i+1:]...)
			break
		}
	}

	for i, l := range h.parent.content {
		if l.Host == h {
			h.parent.content = append(h.parent.content[:i], h.parent.content[i+1:]...)
			break
		}
	}
}
