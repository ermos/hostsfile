package hostsfile

import "runtime"

func linebreak() string {
	return linebreakFromOS(runtime.GOOS)
}

func linebreakFromOS(osName string) string {
	switch osName {
	case "windows":
		return "\r\n"
	default:
		return "\n"
	}
}
