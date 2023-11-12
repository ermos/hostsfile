package hostsfile

import (
	"testing"
)

func TestHostsAddHost(t *testing.T) {
	hosts := &Hosts{
		rows: []hostRow{
			{},
		},
	}

	hosts.AddHost(&Host{})

	if len(hosts.rows) != 2 {
		t.Errorf("AddHost did not add host")
	}
}

func TestHostsAddRaw(t *testing.T) {
	hosts := &Hosts{
		rows: []hostRow{
			{},
		},
	}

	hosts.AddRaw("")

	if len(hosts.rows) != 2 {
		t.Errorf("AddRaw did not add host")
	}
}

func TestRemoveHostsByAddress(t *testing.T) {
	hosts := &Hosts{
		rows: []hostRow{
			{host: &Host{address: "0.0.0.0"}},
			{host: &Host{address: "1.1.1.1"}},
			{host: &Host{address: "0.0.0.0"}},
		},
	}

	hosts.RemoveHostsByAddress("0.0.0.0")

	if len(hosts.rows) != 1 {
		t.Errorf("RemoveHostsByAddress did not remove hosts")
	}
}

func TestRemoveHostsByHostName(t *testing.T) {
	hosts := &Hosts{
		rows: []hostRow{
			{host: &Host{hostNames: []string{"a"}}},
			{host: &Host{hostNames: []string{"b"}}},
			{host: &Host{hostNames: []string{"a"}}},
		},
	}

	hosts.RemoveHostsByHostName("a")

	if len(hosts.rows) != 1 {
		t.Errorf("RemoveHostsByHostName did not remove hosts")
	}
}

func TestRemoveHostsByComment(t *testing.T) {
	hosts := &Hosts{
		rows: []hostRow{
			{host: &Host{comment: "a"}},
			{host: &Host{comment: "b"}},
			{host: &Host{comment: "a"}},
		},
	}

	hosts.RemoveHostsByComment("a")

	if len(hosts.rows) != 1 {
		t.Errorf("RemoveHostsByComment did not remove hosts")
	}
}

func TestRemoveAllHosts(t *testing.T) {
	hosts := &Hosts{
		rows: []hostRow{
			{raw: "a"},
			{host: &Host{address: "0.0.0.0"}},
			{host: &Host{address: "1.1.1.1"}},
			{host: &Host{address: "0.0.0.0"}},
		},
	}

	hosts.RemoveAllHosts()

	if len(hosts.rows) != 1 {
		t.Errorf("RemoveAllHosts did not remove content")
	}
}
