package main

import (
	"testing"
)

func TestCheckIP_OK(t *testing.T) {

   _, isIPGood := CheckIP("127.0.0.1")

	if !isIPGood {
	   	t.Error("Invalid IP")
	}

}

func TestCheckIP_Bad(t *testing.T) {

	_, isIPGood := CheckIP("127.0.1")

	if isIPGood {
		t.Error("IP Is supposed to be Invalid")
	}

}