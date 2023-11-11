package hostsfile

// FindFromHostName returns the host from the hostname.
func (hosts *Hosts) FindFromHostName(hostName string) (host *Host, err error) {
	for _, h := range hosts.hosts {
		for _, d := range h.hostNames {
			if d == hostName {
				return h, nil
			}
		}
	}
	return host, ErrHostNotFound
}

// FindFromAddress returns the first host found based on given address.
func (hosts *Hosts) FindFromAddress(address string) (host *Host, err error) {
	for _, h := range hosts.hosts {
		if h.address == address {
			return h, nil
		}
	}
	return host, ErrHostNotFound
}

// FindAllFromAddress returns all hosts found based on given address.
func (hosts *Hosts) FindAllFromAddress(address string) (hostsFound []*Host) {
	for _, h := range hosts.hosts {
		if h.address == address {
			hostsFound = append(hostsFound, h)
		}
	}
	return hostsFound
}

// FindFromComment returns the first host found based on given comment.
func (hosts *Hosts) FindFromComment(comment string) (host *Host, err error) {
	for _, h := range hosts.hosts {
		if h.comment == comment {
			return h, nil
		}
	}
	return host, ErrHostNotFound
}

// FindAllFromComment returns all hosts found based on given comment.
func (hosts *Hosts) FindAllFromComment(comment string) (hostsFound []*Host) {
	for _, h := range hosts.hosts {
		if h.comment == comment {
			hostsFound = append(hostsFound, h)
		}
	}

	return hostsFound
}
