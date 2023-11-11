package hostsfile

import (
	"errors"
	"os"
	"testing"
)

func TestAddHostWithoutComment(t *testing.T) {
	readFile = func(filename string) ([]byte, error) {
		return addHostRead(), nil
	}

	writeFile = func(filename string, data []byte, perm os.FileMode) error {
		if string(data) != string(expectedAddHostWithoutComment()) {
			t.Errorf("expected `%s`, got `%s`", expectedAddHostWithoutComment(), data)
		}
		return nil
	}

	err := AddHost(Host{
		address:   "ff02::1",
		hostNames: []string{"ip6-allnodes"},
	})
	if err != nil {
		t.Errorf("expected nil, got %s", err)
	}
}

func TestAddHostWithComment(t *testing.T) {
	readFile = func(filename string) ([]byte, error) {
		return addHostRead(), nil
	}

	writeFile = func(filename string, data []byte, perm os.FileMode) error {
		if string(data) != string(expectedAddHostWithComment()) {
			t.Errorf("expected `%s`, got `%s`", expectedAddHostWithComment(), data)
		}
		return nil
	}

	err := AddHost(Host{
		address:   "ff02::1",
		hostNames: []string{"ip6-allnodes"},
		comment:   "this is a comment",
	})
	if err != nil {
		t.Errorf("expected nil, got %s", err)
	}
}

func TestAddHostFromOS(t *testing.T) {
	readFile = func(filename string) ([]byte, error) {
		if filename != "/etc/hosts" {
			t.Errorf("expected `/etc/hosts` for linux, got `%s`", filename)
		}
		return addHostRead(), nil
	}

	writeFile = func(filename string, data []byte, perm os.FileMode) error {
		return nil
	}

	err := AddHostFromOS("linux", Host{
		address:   "ff02::1",
		hostNames: []string{"ip6-allnodes"},
	})
	if err != nil {
		t.Errorf("expected nil, got %s", err)
	}
}

func TestAddHostWithoutAddress(t *testing.T) {
	readFile = func(filename string) ([]byte, error) {
		return addHostRead(), nil
	}

	writeFile = func(filename string, data []byte, perm os.FileMode) error {
		return nil
	}

	err := AddHost(Host{
		hostNames: []string{"ip6-allnodes"},
	})
	if !errors.Is(err, ErrHostAddressIsRequired) {
		t.Errorf("expected `%s`, got `%s`", ErrHostAddressIsRequired, err)
	}
}

func TestAddHostWithoutHostName(t *testing.T) {
	readFile = func(filename string) ([]byte, error) {
		return addHostRead(), nil
	}

	writeFile = func(filename string, data []byte, perm os.FileMode) error {
		return nil
	}

	err := AddHost(Host{
		address: "ff02::1",
	})
	if !errors.Is(err, ErrHostRequireAtLeastOneHostname) {
		t.Errorf("expected `%s`, got `%s`", ErrHostRequireAtLeastOneHostname, err)
	}
}

func TestAddHostInHostFileWithoutLineBreakAtEnd(t *testing.T) {
	readFile = func(filename string) ([]byte, error) {
		return addHostReadWithoutLineBreak(), nil
	}

	writeFile = func(filename string, data []byte, perm os.FileMode) error {
		if string(data) != string(expectedAddHostWithoutComment()) {
			t.Errorf("expected `%s`, got `%s`", expectedAddHostWithoutComment(), data)
		}
		return nil
	}

	err := AddHost(Host{
		address:   "ff02::1",
		hostNames: []string{"ip6-allnodes"},
	})
	if err != nil {
		t.Errorf("expected nil, got %s", err)
	}
}

func TestAddHostWithWrongOSPath(t *testing.T) {
	expectedErr := errors.New("read error")

	readFile = func(filename string) ([]byte, error) {
		return []byte{}, expectedErr
	}

	writeFile = func(filename string, data []byte, perm os.FileMode) error {
		return nil
	}

	err := AddHost(Host{
		address:   "ff02::1",
		hostNames: []string{"ip6-allnodes"},
	})
	if !errors.Is(err, expectedErr) {
		t.Errorf("expected `%s`, got `%s`", expectedErr, err)
	}
}

func TestAddHostWithUnsupportedOperatingSystem(t *testing.T) {
	readFile = func(filename string) ([]byte, error) {
		return []byte{}, nil
	}

	writeFile = func(filename string, data []byte, perm os.FileMode) error {
		return nil
	}

	err := AddHostFromOS("unsupported", Host{
		address:   "ff02::1",
		hostNames: []string{"ip6-allnodes"},
	})
	if !errors.Is(err, ErrUnsupportedOperatingSystem) {
		t.Errorf("expected `%s`, got `%s`", ErrUnsupportedOperatingSystem, err)
	}
}

func addHostRead() []byte {
	return []byte(`
# This is a comment
127.0.0.1	localhost # this is localhost
127.0.1.1	ermos
`)
}

func addHostReadWithoutLineBreak() []byte {
	return []byte(`
# This is a comment
127.0.0.1	localhost # this is localhost
127.0.1.1	ermos`)
}

func expectedAddHostWithoutComment() []byte {
	return []byte(`
# This is a comment
127.0.0.1	localhost # this is localhost
127.0.1.1	ermos
ff02::1 ip6-allnodes
`)
}

func expectedAddHostWithComment() []byte {
	return []byte(`
# This is a comment
127.0.0.1	localhost # this is localhost
127.0.1.1	ermos
ff02::1 ip6-allnodes # this is a comment
`)
}
