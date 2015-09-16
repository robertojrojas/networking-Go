package main

import (
	"net"
	"os"
	"fmt"
)


func CheckIP(ip string) (ipResult net.IP, isGood bool) {

	addr := net.ParseIP(ip)

	ipResult, isGood = addr, addr != nil

	return
}


func main() {

	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s ip-addr\n", os.Args[0])
		os.Exit(1)
	}

	name := os.Args[1]

	addr, isIPGood := CheckIP(name)
	if !isIPGood {
		fmt.Printf("The address %s is Invalid\n", name)
		os.Exit(1)
	}

	fmt.Println("The address is ", addr.String())

	os.Exit(0)

}

