package main

import (
	"github.com/ermos/hostsfile"
)

func main() {
	hosts, err := hostsfile.ParseFromPath("../default.hosts")
	if err != nil {
		panic(err)
	}

	hosts.RemoveAllHosts()

	if err = hosts.FlushToPath("./generated.hosts"); err != nil {
		panic(err)
	}
}
