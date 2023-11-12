package hostsfile

import (
	"runtime"
	"strings"
)

// Parse parses the hosts file based on the current OS.
func Parse() (hosts *Hosts, err error) {
	return ParseFromOS(runtime.GOOS)
}

// ParseFromOS parses the hosts file based on the given OS.
func ParseFromOS(osName string) (hosts *Hosts, err error) {
	path, err := GetSystemPathByOS(osName)
	if err != nil {
		return
	}

	return ParseFromPath(path)
}

// ParseFromPath parses the hosts file based on the given path.
func ParseFromPath(path string) (hosts *Hosts, err error) {
	var contentByte []byte
	var content string

	hosts = &Hosts{}

	contentByte, err = readFile(path)
	if err != nil {
		return
	}

	content = string(contentByte)

	hosts.path = path

main:
	for _, l := range strings.Split(content, linebreak()) {
		var hostNames []string
		var comment string

		fields := strings.Fields(l)

		if len(fields) <= 1 || strings.HasPrefix(fields[0], "#") {
			hosts.rows = append(hosts.rows, hostRow{
				raw: l,
			})
			continue
		}

		hostNames = fields[1:]
		for i, domain := range hostNames {
			if strings.HasPrefix(domain, "#") {
				if i == 0 {
					hosts.rows = append(hosts.rows, hostRow{
						raw: l,
					})
					continue main
				}

				comment = strings.TrimSpace(
					strings.TrimPrefix(
						strings.Join(hostNames[i:], " "),
						"#",
					),
				)
				hostNames = hostNames[:i]

				break
			}
		}

		hosts.rows = append(hosts.rows, hostRow{
			host: &Host{
				address:   fields[0],
				hostNames: hostNames,
				comment:   comment,
				parent:    hosts,
			},
		})
	}

	return hosts, nil
}
