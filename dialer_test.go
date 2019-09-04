package main

import (
	"encoding/xml"
	"net"
	"os"
	"testing"
)

func TestDialUPNP(t *testing.T) {
	target := os.Getenv("UPNP_TEST_TARGET")
	addr := net.ParseIP(target)
	if addr == nil {
		t.Error("wanted ip for UPNP_TEST_TARGET")
	}
	out, err := dialUPNP(addr)
	if err != nil {
		t.Error(err)
	}
	if len(out) == 0 {
		t.Error("got 0 length for output, expected content")
	}
	update := ESATMUpdate{}
	err = xml.Unmarshal(out, &update)
	if err != nil {
		t.Error(err)
	}
	t.Logf("%+v", update)
}
