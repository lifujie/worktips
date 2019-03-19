package main

import (
	"fmt"
)

var (
	hostList = []string{"gitlab.idcos.com", "3Q28132.idcos.com", "www.baidu.com", "www.baidu.comm", "wwww"}
)

func main() {
	for k := range hostList {
		addrs, cname, err := cgoLookupIPCNAME(hostList[k])
		if err != nil {
			fmt.Printf("%s->%s\n", hostList[k], err.Error())
		} else {
			fmt.Printf("%s->%s, %s, %v\n", hostList[k], cname, addrs)
		}
	}
}
