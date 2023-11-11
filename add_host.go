package hostsfile

import (
	"fmt"
	"runtime"
	"strings"
)

// AddHost adds a host to the hosts file based on the current OS.
func AddHost(host Host) error {
	return addHost(runtime.GOOS, host)
}

// AddHostByOS adds a host to the hosts file based on the given OS.
func AddHostByOS(osName string, host Host) error {
	return addHost(osName, host)
}

// addHost adds a host to the hosts file based on the given OS.
func addHost(osName string, host Host) error {
	path, err := GetSystemPathByOS(osName)
	if err != nil {
		return err
	}

	content, err := readFile(path)
	if err != nil {
		return err
	}

	var lineBuilder []string

	if host.Address == "" {
		return ErrHostAddressIsRequired
	}

	lineBuilder = append(lineBuilder, host.Address)

	if len(host.HostNames) == 0 {
		return ErrHostRequireAtLeastOneHostname
	}

	lineBuilder = append(lineBuilder, host.HostNames...)

	if host.Comment != "" {
		lineBuilder = append(lineBuilder, fmt.Sprintf("# %s", host.Comment))
	}

	line := strings.Join(lineBuilder, " ") + "\n"

	if !strings.HasSuffix(string(content), "\n") {
		line = "\n" + line
	}

	content = append(content, []byte(line)...)

	return writeFile(path, content, 0644)
}
