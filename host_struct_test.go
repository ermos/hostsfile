package hostsfile

import "testing"

func TestIsCurrentHost(t *testing.T) {
	host := Host{
		Address: "127.0.0.1",
	}

	if !host.IsCurrentHost() {
		t.Errorf("expected true, got false")
	}
}

func TestIsNotCurrentHost(t *testing.T) {
	host := Host{
		Address: "0.0.0.0",
	}

	if host.IsCurrentHost() {
		t.Errorf("expected false, got true")
	}
}

func TestIsPrivateHost(t *testing.T) {
	host := Host{
		Address: "192.168.1.23",
	}

	if !host.IsPrivateHost() {
		t.Errorf("expected true, got false")
	}
}

func TestIsNotPrivateHost(t *testing.T) {
	host := Host{
		Address: "0.0.0.0",
	}

	if host.IsPrivateHost() {
		t.Errorf("expected false, got true")
	}
}

func TestIsPrivateHostCantParseIPAddress(t *testing.T) {
	host := Host{
		Address: "google.com",
	}

	if host.IsPrivateHost() {
		t.Errorf("expected false, got true")
	}
}

func TestIsPublicHost(t *testing.T) {
	host := Host{
		Address: "0.0.0.0",
	}

	if !host.IsPublicHost() {
		t.Errorf("expected true, got false")
	}
}

func TestIsNotPublicHost(t *testing.T) {
	host := Host{
		Address: "192.168.1.23",
	}

	if host.IsPublicHost() {
		t.Errorf("expected false, got true")
	}
}

func TestHasHostName(t *testing.T) {
	host := Host{
		Address:   "0.0.0.0",
		HostNames: []string{"google.com", "localhost"},
	}

	if !host.HasHostName("localhost") {
		t.Errorf("expected true, got false")
	}
}

func TestHasHostNameNotFound(t *testing.T) {
	host := Host{
		Address:   "0.0.0.0",
		HostNames: []string{"google.com", "localhost"},
	}

	if host.HasHostName("youtube.com") {
		t.Errorf("expected false, got true")
	}
}
