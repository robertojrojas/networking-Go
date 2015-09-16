package main

import (
	"net"
	"os"
	"fmt"
)

func main() {


	// go run <program> google.com:80
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s host:port\n", os.Args[0])
		fmt.Println("Usage: ", os.Args[0], "host:port")
		os.Exit(1)
	}

	addr := os.Args[1]

	// Can use either "tcp", "tcp4" or "tcp6"
	tcpAddr, err := net.ResolveTCPAddr("tcp", addr)
	if err != nil {
		fmt.Println("Resolution error", err.Error())
		os.Exit(1)
	}

	fmt.Printf("Resolved IP %s Port %d\n", tcpAddr.IP, tcpAddr.Port)
	os.Exit(0)

}
