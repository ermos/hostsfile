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
	for i, row := range h.parent.rows {
		if row.host == h {
			h.parent.rows = append(h.parent.rows[:i], h.parent.rows[i+1:]...)
			break
		}
	}
}
