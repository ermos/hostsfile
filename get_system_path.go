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
		path = "/etc/hosts"
		break
	case "windows":
		path = fmt.Sprintf(`%s\System32\drivers\etc\hosts`, os.Getenv("SystemRoot"))
		break
	case "android":
		if _, err = statFile("/system/etc/hosts"); err != nil {
			path = "/etc/hosts"
			break
		}
		path = "/system/etc/hosts"
		break
	case "hurd":
		if _, err = statFile("/hurd/hosts"); err != nil {
			path = "/etc/hosts"
			break
		}
		path = "/hurd/hosts"
		break
	case "illumos":
		path = "/etc/inet/hosts"
		break
	case "plan9":
		path = "/lib/ndb/local"
		break
	case "zos":
		path = "SYS1.PARMLIB(HOSTS)"
		break
	default:
		return "", ErrUnsupportedOperatingSystem
	}

	return path, nil
}
