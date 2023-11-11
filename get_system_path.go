package hostsfile

import (
	"fmt"
	"os"
	"runtime"
)

// GetSystemPath returns the path of the hosts file based on the current OS.
func GetSystemPath() (path string, err error) {
	return getSystemPath(runtime.GOOS)
}

// GetSystemPathByOS returns the path of the hosts file based on the given OS.
func GetSystemPathByOS(osName string) (path string, err error) {
	return getSystemPath(osName)
}

// getSystemPath returns the path of the hosts file based on the given OS.
func getSystemPath(osName string) (path string, err error) {
	switch osName {
	case "linux", "darwin", "aix", "dragonfly", "freebsd", "ios", "netbsd", "openbsd", "solaris":
		return "/etc/hosts", nil
	case "windows":
		return fmt.Sprintf(`%s\System32\drivers\etc\hosts`, os.Getenv("SystemRoot")), nil
	case "android":
		return "/system/etc/hosts", nil // can be "/etc/hosts"
	case "hurd":
		return "/hurd/hosts", nil // can be "/etc/hosts"
	case "illumos":
		return "/etc/inet/hosts", nil
	case "plan9":
		return "/lib/ndb/local", nil
	case "zos":
		return "SYS1.PARMLIB(HOSTS)", nil
	default:
		return "", ErrUnsupportedOperatingSystem
	}
}
