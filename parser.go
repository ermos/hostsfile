package hostsfile

import (
	"runtime"
	"strings"
)

// Parse parses the hosts file based on the current OS.
func Parse() (hosts Hosts, err error) {
	return parseFile(runtime.GOOS)
}

// ParseByOS parses the hosts file based on the given OS.
func ParseByOS(osName string) (hosts Hosts, err error) {
	return parseFile(osName)
}

// parseFile parses the hosts file based on the given OS.
func parseFile(osName string) (hosts Hosts, err error) {
	var content []byte

	path, err := GetSystemPathByOS(osName)
	if err != nil {
		return nil, err
	}

	content, err = readFile(path)
	if err != nil {
		return nil, err
	}

	return parseContent(string(content))
}

// parseContent parses the content of the hosts file.
func parseContent(content string) (hosts []Host, err error) {
main:
	for _, line := range strings.Split(content, "\n") {
		var hostNames []string
		var comment string

		fields := strings.Fields(line)

		if len(fields) <= 1 || strings.HasPrefix(fields[0], "#") {
			continue
		}

		hostNames = fields[1:]
		for i, domain := range hostNames {
			if strings.HasPrefix(domain, "#") {
				if i == 0 {
					continue main
				}

				hostNames = hostNames[:i]
				comment = strings.TrimSpace(
					strings.TrimPrefix(
						strings.Join(hostNames[i:], " "),
						"#",
					),
				)
				break
			}
		}

		hosts = append(hosts, Host{
			Address:   fields[0],
			HostNames: hostNames,
			Comment:   comment,
		})
	}

	return hosts, nil
}
