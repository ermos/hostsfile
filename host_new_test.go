package hostsfile

import "testing"

func TestNewHost(t *testing.T) {
	host := NewHost("0.0.0.0", []string{"localhost"})

	if host.address != "0.0.0.0" {
		t.Errorf("NewHost did not set address")
	}

	if host.hostNames[0] != "localhost" {
		t.Errorf("NewHost did not set hostname")
	}

	if host.comment != "" {
		t.Errorf("NewHost did not set comment")
	}
}

func TestNewHostWithComment(t *testing.T) {
	host := NewHost("0.0.0.0", []string{"localhost"}, WithComment("this is a comment"))

	if host.comment != "this is a comment" {
		t.Errorf("NewHost did not set comment")
	}
}
