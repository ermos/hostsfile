package hostsfile

import (
	"errors"
	"os"
	"testing"
)

func TestHostsAddHost(t *testing.T) {
	hosts := &Hosts{
		hosts: []*Host{
			{},
		},
	}

	hosts.AddHost(&Host{})

	if len(hosts.hosts) != 2 {
		t.Errorf("AddHost did not add host")
	}
}

func TestRemoveHostsByAddress(t *testing.T) {
	hosts := &Hosts{
		hosts: []*Host{
			{address: "0.0.0.0"},
			{address: "1.1.1.1"},
			{address: "0.0.0.0"},
		},
	}

	hosts.RemoveHostsByAddress("0.0.0.0")

	if len(hosts.hosts) != 1 {
		t.Errorf("RemoveHostsByAddress did not remove hosts")
	}
}

func TestRemoveHostsByHostName(t *testing.T) {
	hosts := &Hosts{
		hosts: []*Host{
			{hostNames: []string{"a"}},
			{hostNames: []string{"b"}},
			{hostNames: []string{"a"}},
		},
	}

	hosts.RemoveHostsByHostName("a")

	if len(hosts.hosts) != 1 {
		t.Errorf("RemoveHostsByHostName did not remove hosts")
	}
}

func TestRemoveHostsByComment(t *testing.T) {
	hosts := &Hosts{
		hosts: []*Host{
			{comment: "a"},
			{comment: "b"},
			{comment: "a"},
		},
	}

	hosts.RemoveHostsByComment("a")

	if len(hosts.hosts) != 1 {
		t.Errorf("RemoveHostsByComment did not remove hosts")
	}
}

func TestRemoveAllHosts(t *testing.T) {
	hosts := &Hosts{
		hosts: []*Host{
			{address: "0.0.0.0"},
			{address: "1.1.1.1"},
			{address: "0.0.0.0"},
		},
		content: []line{
			{IsHost: false},
			{IsHost: true},
			{IsHost: true},
			{IsHost: true},
		},
	}

	hosts.RemoveAllHosts()

	if len(hosts.hosts) != 0 {
		t.Errorf("RemoveAllHosts did not remove hosts")
	}

	if len(hosts.content) != 1 {
		t.Errorf("RemoveAllHosts did not remove content")
	}
}

func TestFlush(t *testing.T) {
	hosts := &Hosts{
		content: []line{
			{IsHost: false, Content: "# test"},
			{IsHost: true, Host: &Host{
				address:   "1.1.1.1",
				hostNames: []string{"a", "b"},
				comment:   "c",
			}},
			{IsHost: true, Content: "not used"},
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

	if len(hosts.content) != 3 {
		t.Errorf("Flush did not keep content")
	}
}

func TestFlushWithInvalidHost(t *testing.T) {
	hosts := &Hosts{
		content: []line{
			{IsHost: false, Content: "# test"},
			{IsHost: true, Host: &Host{
				address:   "1.1.1.1",
				hostNames: []string{"a", "b"},
				comment:   "c",
			}},
			{IsHost: true, Host: &Host{
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
