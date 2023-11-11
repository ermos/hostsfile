package hostsfile

import "testing"

func TestGetAddress(t *testing.T) {
	h := &Host{
		address: "0.0.0.0",
	}

	if h.GetAddress() != "0.0.0.0" {
		t.Errorf("GetAddress returned incorrect address")
	}
}

func TestGetHostNames(t *testing.T) {
	h := &Host{
		hostNames: []string{"localhost"},
	}

	if h.GetHostNames()[0] != "localhost" {
		t.Errorf("GetHostNames returned incorrect hostnames")
	}
}

func TestGetComment(t *testing.T) {
	h := &Host{
		comment: "comment",
	}

	if h.GetComment() != "comment" {
		t.Errorf("GetComment returned incorrect comment")
	}
}
