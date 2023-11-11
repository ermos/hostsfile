package hostsfile

import "testing"

func TestGetPath(t *testing.T) {
	hosts := &Hosts{
		path: "path",
	}

	if hosts.GetPath() != "path" {
		t.Errorf("GetPath returned incorrect path")
	}
}

func TestGetHosts(t *testing.T) {
	hosts := &Hosts{
		hosts: []*Host{
			{},
		},
	}

	if len(hosts.GetHosts()) != 1 {
		t.Errorf("GetHosts returned incorrect hosts")
	}
}
