/*
LookupHost
 */

package main

import (
	"net"
	"os"
	"fmt"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s hostname\n", os.Args[0])
		os.Exit(1)
	}
	name := os.Args[1]

	/*
		One of these addresses will be labelled as the "canonical" host name.
		If you wish to find the canonical name,
		use func LookupCNAME(name string) (cname string, err os.Error)
	*/
	addrs, err := net.LookupHost(name)
	if err != nil {
		fmt.Println("Error: ", err.Error())
		os.Exit(2)
	}

	for _, s := range addrs {

		if strings.Contains(s, ":") {
			fmt.Printf("IPV6 %s\n", s)
		} else {
			fmt.Printf("IPV4 %s\n", s)
		}


	}
	os.Exit(0)
}