package main

import "C"

import (
	"net"
)

//go:linkname cgoLookupIPCNAME net.cgoLookupIPCNAME
func cgoLookupIPCNAME(name string) (addrs []net.IPAddr, cname string, err error)
