package main

import (
	"github.com/ermos/hostsfile"
)

func main() {
	hosts := hostsfile.NewHosts("./generated.hosts")

	hosts.AddHost(hostsfile.NewHost(
		"1.1.1.1",
		[]string{"cloudflare.com"},
		hostsfile.WithComment("Cloudflare DNS"),
	))

	if err := hosts.Flush(); err != nil {
		panic(err)
	}
}
