package hostsfile

import (
	"errors"
	"testing"
)

func TestParse(t *testing.T) {
	readFile = func(filename string) ([]byte, error) {
		return parseRead(), nil
	}

	hosts, err := Parse()
	if err != nil {
		t.Errorf("expected nil, got %s", err)
	}

	if len(hosts) != 2 {
		t.Errorf("expected 2 hosts, got %d", len(hosts))
	}
}

func TestParseByOS(t *testing.T) {
	readFile = func(filename string) ([]byte, error) {
		return parseRead(), nil
	}

	hosts, err := ParseByOS("linux")
	if err != nil {
		t.Errorf("expected nil, got %s", err)
	}

	if len(hosts) != 2 {
		t.Errorf("expected 2 hosts, got %d", len(hosts))
	}
}

func TestParseWithWrongOSPath(t *testing.T) {
	expectedErr := errors.New("read error")

	readFile = func(filename string) ([]byte, error) {
		return []byte{}, expectedErr
	}

	_, err := Parse()
	if !errors.Is(err, expectedErr) {
		t.Errorf("expected `%s`, got `%s`", expectedErr, err)
	}
}

func TestParseWithUnsupportedOperatingSystem(t *testing.T) {
	readFile = func(filename string) ([]byte, error) {
		return []byte{}, nil
	}

	_, err := ParseByOS("unsupported")
	if !errors.Is(err, ErrUnsupportedOperatingSystem) {
		t.Errorf("expected `%s`, got `%s`", ErrUnsupportedOperatingSystem, err)
	}
}

func parseRead() []byte {
	return []byte(`
# This is a comment
# 127.0.0.2 cat
127.0.0.1	localhost # this is localhost
127.0.1.1	ermos
0.0.0.0
1.1.1.1 # wrong
2.2.2.2 #wrong
`)
}
