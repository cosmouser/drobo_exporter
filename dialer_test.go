package main

import (
	"encoding/xml"
	"os"
	"testing"
)

func TestDialUPNP(t *testing.T) {
	target := os.Getenv("UPNP_TEST_TARGET")
	out, err := dialUPNP(target)
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
