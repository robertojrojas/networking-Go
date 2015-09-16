package main

import (
	"net"
	"os"
	"fmt"
	"io/ioutil"
)

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s\n", err.Error())
		os.Exit(1)
	}
}

func main() {

	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s host:port\n", os.Args[0])
		os.Exit(1)
	}

	service := os.Args[1]

	// There may be syntax errors in the address specified
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkError(err)

	// The attempt to connect to the remote service may fail.
	// For example, the service requested might not be running,
	// or there may be no such host connected to the network
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	checkError(err)

	// Although a connection has been established,
	// writes to the service might fail if the connection has died suddenly,
	// or the network times out
	_, err = conn.Write([]byte("HEAD / HTTP/1.0\r\n\r\n"))
	checkError(err)

	// Similarly, the reads might fail
	result, err := ioutil.ReadAll(conn)
	checkError(err)

	fmt.Println(string(result))

	os.Exit(0)
}
