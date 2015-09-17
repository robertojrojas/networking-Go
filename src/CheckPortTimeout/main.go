package main

import (
	"net"
	"os"
	"fmt"
	"time"
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

func IsPortOpenWithTimeout(isOpenCh chan bool, connectionType, addr, timeout string) {

	duration, _   := time.ParseDuration(timeout)
	sleepInterval := 100 * time.Millisecond
	theTimer      := time.NewTimer(duration)

	for {
		select {
			case <-theTimer.C:
				isOpenCh <- false
			default :
				if (IsPortOpen(connectionType, addr)) {
					isOpenCh <- true
					return
				}

				time.Sleep(sleepInterval)
		}
	}

}

func main() {

	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s host:port \n", os.Args[0])
		os.Exit(1)
	}

	service := os.Args[1]

	isOpenCh := make(chan bool)
	go IsPortOpenWithTimeout(isOpenCh, "tcp", service, "5s")

	fmt.Printf("Waiting ...\n")
	isOpenResult := <-isOpenCh
	fmt.Printf("Is Port Open? %v\n", isOpenResult)

}
