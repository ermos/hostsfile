package hostsfile

import (
	"errors"
	"testing"
)

func TestFindFromHostName(t *testing.T) {
	host, err := fakeHosts().FindFromHostName("google.com")
	if err != nil {
		t.Errorf("expected nil, got %s", err)
	}

	if host.address != "0.0.0.0" {
		t.Errorf("wrong host address, expected `0.0.0.0`, got `%s`", host.address)
	}
}

func TestCannotFindFromHostName(t *testing.T) {
	_, err := fakeHosts().FindFromHostName("smiti.fr")
	if !errors.Is(err, ErrHostNotFound) {
		t.Errorf("expected `%s`, got `%s`", ErrHostNotFound, err)
	}
}

func TestFindFromAddress(t *testing.T) {
	host, err := fakeHosts().FindFromAddress("1.1.1.1")
	if err != nil {
		t.Errorf("expected nil, got %s", err)
	}

	if host.hostNames[0] != "yahoo.com" {
		t.Errorf("wrong hostname, expected `yahoo.com, got `%s`", host.hostNames[0])
	}
}

func TestCannotFindFromAddress(t *testing.T) {
	_, err := fakeHosts().FindFromAddress("9.9.9.9")
	if !errors.Is(err, ErrHostNotFound) {
		t.Errorf("expected `%s`, got `%s`", ErrHostNotFound, err)
	}
}

func TestFindAllFromAddress(t *testing.T) {
	hosts := fakeHosts().FindAllFromAddress("1.1.1.1")
	if len(hosts) != 2 {
		t.Errorf("wrong hosts length, expected `2`, got `%d`", len(hosts))
	}

	if hosts[0].hostNames[0] != "yahoo.com" && hosts[1].hostNames[0] != "youtube.com" {
		t.Errorf(
			"wrong hosts address, expected `yahoo.com` and `youtube.com`, got `%s` and `%s`",
			hosts[0].hostNames[0],
			hosts[1].hostNames[0],
		)
	}
}

func TestFindFromComment(t *testing.T) {
	host, err := fakeHosts().FindFromComment("this is a comment")
	if err != nil {
		t.Errorf("expected nil, got %s", err)
	}

	if host.address != "1.1.1.1" {
		t.Errorf("wrong host address, expected `1.1.1.1`, got `%s`", host.address)
	}
}

func TestCannotFindFromComment(t *testing.T) {
	_, err := fakeHosts().FindFromComment("comment not found")
	if !errors.Is(err, ErrHostNotFound) {
		t.Errorf("expected `%s`, got `%s`", ErrHostNotFound, err)
	}
}

func TestFindAllFromComment(t *testing.T) {
	hosts := fakeHosts().FindAllFromComment("this is a comment")
	if len(hosts) != 2 {
		t.Errorf("wrong hosts length, expected `2`, got `%d`", len(hosts))
	}

	if hosts[0].address != "1.1.1.1" && hosts[1].address != "2.2.2.2" {
		t.Errorf(
			"wrong hosts address, expected `1.1.1.1` and `2.2.2.2`, got `%s` and `%s`",
			hosts[0].address,
			hosts[1].address,
		)
	}
}

func fakeHosts() *Hosts {
	return &Hosts{
		hosts: []*Host{
			{
				address:   "0.0.0.0",
				hostNames: []string{"google.com", "localhost"},
			},
			{
				address:   "1.1.1.1",
				hostNames: []string{"yahoo.com"},
				comment:   "this is a comment",
			},
			{
				address:   "2.2.2.2",
				hostNames: []string{"google.com", "localhost"},
				comment:   "this is a comment",
			},
			{
				address:   "1.1.1.1",
				hostNames: []string{"youtube.com"},
			},
		},
	}
}
