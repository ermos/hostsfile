package main

import (
	"github.com/ermos/hostsfile"
)

func main() {
	hosts, err := hostsfile.ParseFromPath("../default.hosts")
	if err != nil {
		panic(err)
	}

	hosts.AddHost(hostsfile.NewHost(
		"1.1.1.1",
		[]string{"cloudflare.com"},
		hostsfile.WithComment("Cloudflare DNS"),
	))

	if err = hosts.FlushToPath("./generated.hosts"); err != nil {
		panic(err)
	}
}
