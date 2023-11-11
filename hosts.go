package hostsfile

// Hosts represents a hosts file.
type Hosts struct {
	path    string
	hosts   []*Host
	content []line
}

type line struct {
	Content string
	Host    *Host
	IsHost  bool
}

func (hosts *Hosts) GetPath() string {
	return hosts.path
}

func (hosts *Hosts) GetHosts() []*Host {
	return hosts.hosts
}
