package hostsfile

import (
	"errors"
	"os"
	"testing"
)

func TestFlush(t *testing.T) {
	hosts := &Hosts{
		rows: []hostRow{
			{raw: "# test"},
			{host: &Host{
				address:   "1.1.1.1",
				hostNames: []string{"a", "b"},
				comment:   "c",
			}},
			{},
		},
	}

	writeFile = func(filename string, data []byte, perm os.FileMode) error {
		if string(data) != "# test\n1.1.1.1 a b # c\n" {
			t.Errorf("Flush did not write correct data : `%s`, expected `%s`", string(data), "# test\n1.1.1.1 a b # c\n")
		}
		return nil
	}

	err := hosts.Flush()
	if err != nil {
		t.Errorf("Flush returned error: %v", err)
	}

	if len(hosts.rows) != 3 {
		t.Errorf("Flush did not keep content")
	}
}

func TestFlushWithInvalidHost(t *testing.T) {
	hosts := &Hosts{
		rows: []hostRow{
			{raw: "# test"},
			{host: &Host{
				address:   "1.1.1.1",
				hostNames: []string{"a", "b"},
				comment:   "c",
			}},
			{host: &Host{
				address: "2.2.2.2",
				comment: "c",
			}},
		},
	}

	writeFile = func(filename string, data []byte, perm os.FileMode) error {
		return nil
	}

	err := hosts.Flush()
	if !errors.Is(err, ErrHostRequireAtLeastOneHostname) {
		t.Errorf("Flush did not return correct error: %v", err)
	}
}
