package hostsfile

import (
	"os"
	"testing"
)

func TestTestFakeReadFile(t *testing.T) {
	var isPassed bool

	TestFakeReadFile(func(name string) ([]byte, error) {
		isPassed = true
		return []byte{}, nil
	})

	if _, err := readFile("test"); err != nil {
		t.Errorf("TestFakeReadFile returned error: %v", err)
	}

	if !isPassed {
		t.Errorf("TestFakeReadFile did not pass")
	}
}

func TestTestFakeWriteFile(t *testing.T) {
	var isPassed bool

	TestFakeWriteFile(func(name string, data []byte, perm os.FileMode) error {
		isPassed = true
		return nil
	})

	if err := writeFile("test", []byte{}, 0); err != nil {
		t.Errorf("TestFakeWriteFile returned error: %v", err)
	}

	if !isPassed {
		t.Errorf("TestFakeWriteFile did not pass")
	}
}

func TestTestFakeStatFile(t *testing.T) {
	var isPassed bool

	TestFakeStatFile(func(name string) (os.FileInfo, error) {
		isPassed = true
		return nil, nil
	})

	if _, err := statFile("test"); err != nil {
		t.Errorf("TestFakeStatFile returned error: %v", err)
	}

	if !isPassed {
		t.Errorf("TestFakeStatFile did not pass")
	}
}
