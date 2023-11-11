package hostsfile

import (
	"strings"
)

// AddHost adds a host to the hosts file.
func (hosts *Hosts) AddHost(host *Host) {
	hosts.hosts = append(hosts.hosts, host)
}

// RemoveHostsByAddress removes all hosts with the given address.
func (hosts *Hosts) RemoveHostsByAddress(address string) {
	var updatedHosts []*Host

	for _, h := range hosts.hosts {
		if h.address != address {
			updatedHosts = append(updatedHosts, h)
		}
	}

	hosts.hosts = updatedHosts
}

// RemoveHostsByHostName removes all hosts with the given hostname.
func (hosts *Hosts) RemoveHostsByHostName(hostName string) {
	var updatedHosts []*Host

	for _, h := range hosts.hosts {
		keep := true

		for _, d := range h.hostNames {
			if d == hostName {
				keep = false
				break
			}
		}

		if keep {
			updatedHosts = append(updatedHosts, h)
		}
	}

	hosts.hosts = updatedHosts
}

// RemoveHostsByComment removes all hosts with the given comment.
func (hosts *Hosts) RemoveHostsByComment(comment string) {
	var updatedHosts []*Host

	for _, h := range hosts.hosts {
		if h.comment != comment {
			updatedHosts = append(updatedHosts, h)
		}
	}

	hosts.hosts = updatedHosts
}

// RemoveAllHosts removes all hosts.
func (hosts *Hosts) RemoveAllHosts() {
	hosts.hosts = nil

	var content []line

	for _, l := range hosts.content {
		if !l.IsHost {
			content = append(content, l)
		}
	}

	hosts.content = content
}

// Flush writes the hosts file to disk.
func (hosts *Hosts) Flush() error {
	var content []string

	for _, l := range hosts.content {
		if !l.IsHost {
			content = append(content, l.Content)
			continue
		}

		if l.Host == nil {
			continue
		}

		c, err := l.Host.ToString()
		if err != nil {
			return err
		}

		content = append(content, c)
	}

	return writeFile(hosts.path, []byte(strings.Join(content, linebreak())+linebreak()), 0644)
}
