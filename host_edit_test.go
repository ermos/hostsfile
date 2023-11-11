package hostsfile

import "testing"

func TestAddHost(t *testing.T) {
	h := &Host{
		hostNames: []string{"localhost"},
	}

	h.AddHostName("example.com")

	if len(h.hostNames) != 2 {
		t.Errorf("AddHostName did not add hostname")
	}
}

func TestAddHostHandleDuplicate(t *testing.T) {
	h := &Host{
		hostNames: []string{"localhost"},
	}

	h.AddHostName("localhost")

	if len(h.hostNames) != 1 {
		t.Errorf("AddHostName added duplicate hostname")
	}
}

func TestRemoveHost(t *testing.T) {
	h := &Host{
		hostNames: []string{"localhost"},
	}

	h.RemoveHostName("localhost")

	if len(h.hostNames) != 0 {
		t.Errorf("RemoveHostName did not remove hostname")
	}
}

func TestRemoveHostHandleNonExistent(t *testing.T) {
	h := &Host{
		hostNames: []string{"localhost"},
	}

	h.RemoveHostName("example.com")

	if len(h.hostNames) != 1 {
		t.Errorf("RemoveHostName removed non-existent hostname")
	}
}

func TestSetComment(t *testing.T) {
	h := &Host{
		comment: "comment",
	}

	h.SetComment("new comment")

	if h.comment != "new comment" {
		t.Errorf("SetComment did not set comment")
	}
}

func TestSetAddress(t *testing.T) {
	h := &Host{
		address: "0.0.0.0",
	}

	h.SetAddress("1.1.1.1")

	if h.address != "1.1.1.1" {
		t.Errorf("SetAddress did not set address")
	}
}

func TestRemove(t *testing.T) {
	host := &Host{
		hostNames: []string{"localhost"},
		address:   "0.0.0.0",
	}

	hosts := &Hosts{
		hosts: []*Host{host},
		content: []line{
			{
				IsHost: true,
				Host:   host,
			},
		},
	}

	host.parent = hosts

	host.Remove()

	if len(hosts.hosts) != 0 {
		t.Errorf("Remove did not remove host")
	}

	if len(hosts.content) != 0 {
		t.Errorf("Remove did not remove host from content")
	}
}
