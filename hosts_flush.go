package hostsfile

import "strings"

// Flush writes the hosts file to disk.
func (hosts *Hosts) Flush() error {
	return hosts.FlushToPath(hosts.path)
}

// FlushToPath writes the hosts file to the given path.
func (hosts *Hosts) FlushToPath(path string) error {
	var content []string

	for _, l := range hosts.rows {
		if l.host == nil {
			content = append(content, l.raw)
			continue
		}

		c, err := l.host.ToString()
		if err != nil {
			return err
		}

		content = append(content, c)
	}

	return writeFile(path, []byte(strings.Join(content, linebreak())), 0644)
}
