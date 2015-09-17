package main

import (
	"net"
	"os"
	"fmt"
)

func IsPortOpen(connectionType, addr string) bool {

	tcpAddr, err := net.ResolveTCPAddr(connectionType, addr)
	if err == nil {
		_, err = net.DialTCP(connectionType, nil, tcpAddr)
		if err != nil {
			return false
		} else {
			return true
		}
	}

	return false
}

func main() {

	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s host:port\n", os.Args[0])
		os.Exit(1)
	}

	service := os.Args[1]
	fmt.Printf("Is Port Open? %v\n", IsPortOpen("tcp4", service))


}
