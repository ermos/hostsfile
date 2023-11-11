package main

import (
	"fmt"
	"github.com/ermos/hostsfile"
)

func main() {
	hosts, err := hostsfile.ParseFromPath("./m/hosts")
	if err != nil {
		panic(err)
	}

	hosts.RemoveAllHosts()

	fmt.Println(hosts.Flush())
}
