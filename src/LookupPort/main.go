package main

import (
	"net"
	"os"
	"fmt"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s service [network-type] \n", os.Args[0])
		os.Exit(1)
	}

	service := os.Args[1]

	var networkType string

	if len(os.Args) == 3 {
		networkType = os.Args[2]

		port, err := net.LookupPort(networkType, service)
		if err != nil {
			fmt.Println("Error: ", err.Error())
			os.Exit(2)
		}

		fmt.Printf("Service %s network Type %s port %d\n", service, networkType, port)

	} else {

		networkTypes := []string {"tcp", "udp"}

		var output string

		for _, networkType := range networkTypes {
			port, err := net.LookupPort(networkType, service)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error: %s\n", err.Error())
				continue
			} else {
				output = fmt.Sprintf("%s %s/%d", output, networkType, port)
			}
		}

		if len(output) > 0 {
			fmt.Printf("Service %s port(s): %s\n", service, output)
		}

	}


	os.Exit(0)
}
