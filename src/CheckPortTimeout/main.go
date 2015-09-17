package main

import (
	"net"
	"os"
	"fmt"
	"time"
	"log"
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

	duration, durationParseErr   := time.ParseDuration(timeout)
	if durationParseErr != nil {
		log.Fatal(durationParseErr)
	}
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

	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s host:port [timeout - 100ms, 5s, 1m]\n", os.Args[0])
		os.Exit(1)
	}
	service := os.Args[1]

	duration := "5s"
	if len(os.Args) > 2 {
		duration = os.Args[2]
	}

	fmt.Printf("Timeout %v\n", duration)

	isOpenCh := make(chan bool)
	go IsPortOpenWithTimeout(isOpenCh, "tcp", service, duration)

	fmt.Printf("Check if %s is listening ...", service)
	isOpenResult := <-isOpenCh
	fmt.Printf("Port Open? %v\n", isOpenResult)

}
