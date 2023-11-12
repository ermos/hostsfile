package main

import (
	"github.com/ermos/hostsfile"
)

func main() {
	hosts, err := hostsfile.LoadFromPath("../default.hosts")
	if err != nil {
		panic(err)
	}

	hosts.AddRaw("# Cloudfare DNS")

	hosts.AddHost(hostsfile.NewHost(
		"1.1.1.1",
		[]string{"cloudflare.com"},
	))

	if err = hosts.FlushToPath("./generated.hosts"); err != nil {
		panic(err)
	}
}
