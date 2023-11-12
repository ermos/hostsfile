package main

import (
	"github.com/ermos/hostsfile"
)

func main() {
	hosts, err := hostsfile.LoadFromPath("../default.hosts")
	if err != nil {
		panic(err)
	}

	host, err := hosts.FindFromAddress("127.0.0.2")
	if err != nil {
		panic(err)
	}

	host.SetComment("This is a comment")
	host.AddHostName("world2.lc")

	if err = hosts.FlushToPath("./generated.hosts"); err != nil {
		panic(err)
	}
}
