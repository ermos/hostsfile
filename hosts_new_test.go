package hostsfile

import "testing"

func TestNewHosts(t *testing.T) {
	hosts := NewHosts("/etc/hosts")

	if hosts.path != "/etc/hosts" {
		t.Errorf("expected `/etc/hosts`, got `%s`", hosts.path)
	}
}
