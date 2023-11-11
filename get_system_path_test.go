package hostsfile

import (
	"errors"
	"os"
	"runtime"
	"testing"
)

func TestGetSystemPath(t *testing.T) {
	p1, err := GetSystemPath()
	if err != nil {
		t.Errorf("GetSystemPath returned an error: %v", err)
	}

	p2, err := GetSystemPathByOS(runtime.GOOS)
	if err != nil {
		t.Errorf("GetSystemPathByOS returned an error: %v", err)
	}

	if p1 != p2 {
		t.Errorf("GetSystemPath returned incorrect path")
	}
}

func TestGetSystemPathByOS(t *testing.T) {
	err := os.Setenv("SystemRoot", `C:\Windows`)
	if err != nil {
		t.Errorf("Setenv returned an error: %v", err)
	}

	testCases := map[string]string{
		"linux":     "/etc/hosts",
		"darwin":    "/etc/hosts",
		"aix":       "/etc/hosts",
		"dragonfly": "/etc/hosts",
		"freebsd":   "/etc/hosts",
		"ios":       "/etc/hosts",
		"netbsd":    "/etc/hosts",
		"openbsd":   "/etc/hosts",
		"solaris":   "/etc/hosts",
		"windows":   `C:\Windows\System32\drivers\etc\hosts`,
		"android":   "/system/etc/hosts",
		"hurd":      "/hurd/hosts",
		"illumos":   "/etc/inet/hosts",
		"plan9":     "/lib/ndb/local",
		"zos":       "SYS1.PARMLIB(HOSTS)",
	}

	statFile = func(path string) (os.FileInfo, error) {
		return nil, nil
	}

	for osName, expectedPath := range testCases {
		var path string

		path, err = GetSystemPathByOS(osName)
		if err != nil {
			t.Errorf("GetSystemPathByOS for %s returned an error: %v", osName, err)
		}

		if path != expectedPath {
			t.Errorf("GetSystemPathByOS for %s returned incorrect path: got %s, want %s", osName, path, expectedPath)
		}
	}
}

func TestGetSystemPathByOSWithInvalidOS(t *testing.T) {
	_, err := GetSystemPathByOS("invalid")
	if !errors.Is(err, ErrUnsupportedOperatingSystem) {
		t.Errorf("expected `%s`, got `%s`", ErrUnsupportedOperatingSystem, err)
	}
}

func TestGetSystemPathByOSSpecialCases(t *testing.T) {
	testCases := map[string]string{
		"android": "/etc/hosts",
		"hurd":    "/etc/hosts",
	}

	statFile = func(path string) (os.FileInfo, error) {
		return nil, errors.New("file not found")
	}

	for osName, expectedPath := range testCases {
		var path string

		path, err := GetSystemPathByOS(osName)
		if err != nil {
			t.Errorf("GetSystemPathByOS for %s returned an error: %v", osName, err)
		}

		if path != expectedPath {
			t.Errorf("GetSystemPathByOS for %s returned incorrect path: got %s, want %s", osName, path, expectedPath)
		}
	}
}
