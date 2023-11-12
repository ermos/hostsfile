package hostsfile

// FindFromHostName returns the host from the hostname.
func (hosts *Hosts) FindFromHostName(hostName string) (*Host, error) {
	for _, row := range hosts.rows {
		if row.host == nil {
			continue
		}

		for _, d := range row.host.hostNames {
			if d == hostName {
				return row.host, nil
			}
		}
	}

	return nil, ErrHostNotFound
}

// FindFromAddress returns the first host found based on given address.
func (hosts *Hosts) FindFromAddress(address string) (*Host, error) {
	for _, row := range hosts.rows {
		if row.host != nil && row.host.address == address {
			return row.host, nil
		}
	}

	return nil, ErrHostNotFound
}

// FindAllFromAddress returns all hosts found based on given address.
func (hosts *Hosts) FindAllFromAddress(address string) []*Host {
	var list []*Host

	for _, row := range hosts.rows {
		if row.host != nil && row.host.address == address {
			list = append(list, row.host)
		}
	}

	return list
}

// FindFromComment returns the first host found based on given comment.
func (hosts *Hosts) FindFromComment(comment string) (*Host, error) {
	for _, row := range hosts.rows {
		if row.host != nil && row.host.comment == comment {
			return row.host, nil
		}
	}

	return nil, ErrHostNotFound
}

// FindAllFromComment returns all hosts found based on given comment.
func (hosts *Hosts) FindAllFromComment(comment string) []*Host {
	var list []*Host

	for _, row := range hosts.rows {
		if row.host != nil && row.host.comment == comment {
			list = append(list, row.host)
		}
	}

	return list
}
