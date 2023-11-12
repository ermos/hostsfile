package hostsfile

// AddHost adds a host to the hosts file.
func (hosts *Hosts) AddHost(host *Host) {
	hosts.rows = append(hosts.rows, hostRow{host: host})
}

func (hosts *Hosts) AddRaw(raw string) {
	hosts.rows = append(hosts.rows, hostRow{raw: raw})
}

// RemoveHostsByAddress removes all hosts with the given address.
func (hosts *Hosts) RemoveHostsByAddress(address string) {
	var rows []hostRow

	for _, row := range hosts.rows {
		if row.host == nil || row.host.address != address {
			rows = append(rows, row)
		}
	}

	hosts.rows = rows
}

// RemoveHostsByHostName removes all hosts with the given hostname.
func (hosts *Hosts) RemoveHostsByHostName(hostName string) {
	var rows []hostRow

	for _, row := range hosts.rows {
		keep := true

		if row.host != nil {
			for _, d := range row.host.hostNames {
				if d == hostName {
					keep = false
					break
				}
			}
		}

		if keep {
			rows = append(rows, row)
		}
	}

	hosts.rows = rows
}

// RemoveHostsByComment removes all hosts with the given comment.
func (hosts *Hosts) RemoveHostsByComment(comment string) {
	var rows []hostRow

	for _, row := range hosts.rows {
		if row.host == nil || row.host.comment != comment {
			rows = append(rows, row)
		}
	}

	hosts.rows = rows
}

// RemoveAllHosts removes all hosts.
func (hosts *Hosts) RemoveAllHosts() {
	var rows []hostRow

	for _, row := range hosts.rows {
		if row.host == nil && row.raw != "" {
			rows = append(rows, row)
		}
	}

	hosts.rows = rows
}
