package hostsfile

// Hosts represents a hosts file.
type Hosts struct {
	path string
	rows []hostRow
}

type hostRow struct {
	raw  string
	host *Host
}

func (hosts *Hosts) GetPath() string {
	return hosts.path
}

func (hosts *Hosts) GetHosts() []*Host {
	var list []*Host

	for _, l := range hosts.rows {
		if l.host != nil {
			list = append(list, l.host)
		}
	}

	return list
}
