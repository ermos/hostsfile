package hostsfile

import (
	"runtime"
	"strings"
)

// AddHost adds a host to the hosts file based on the current OS.
func AddHost(host Host) error {
	return AddHostFromOS(runtime.GOOS, host)
}

// AddHostFromOS adds a host to the hosts file based on the given OS.
func AddHostFromOS(osName string, host Host) error {
	path, err := GetSystemPathByOS(osName)
	if err != nil {
		return err
	}

	return AddHostFromPath(path, host)
}

// AddHostFromPath adds a host to the hosts file based on the given path.
func AddHostFromPath(path string, host Host) error {
	content, err := readFile(path)
	if err != nil {
		return err
	}

	l, err := host.ToString()
	if err != nil {
		return err
	}

	l += linebreak()

	if !strings.HasSuffix(string(content), linebreak()) {
		l = linebreak() + l
	}

	content = append(content, []byte(l)...)

	return writeFile(path, content, 0644)
}
