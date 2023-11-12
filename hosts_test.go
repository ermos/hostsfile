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
		rows: []hostRow{
			{raw: "# comment"},
			{host: &Host{}},
			{host: &Host{}},
			{host: &Host{}},
			{raw: "# comment"},
		},
	}

	if len(hosts.GetHosts()) != 3 {
		t.Errorf("GetHosts returned incorrect hosts")
	}
}
