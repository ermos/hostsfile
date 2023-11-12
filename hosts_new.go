package hostsfile

func NewHosts(path string) *Hosts {
	return &Hosts{
		path: path,
	}
}
