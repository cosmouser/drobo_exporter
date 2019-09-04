package main

import (
	"bufio"
	"bytes"
	"net"
	"strings"
	"time"
)

const (
	uPNPPort     = 5000
	nasdXMLOpen  = "<ESATMUpdate>"
	nasdXMLClose = "</ESATMUpdate>"
)

func dialUPNP(ip net.IP) (out []byte, err error) {
	addr := net.TCPAddr{
		IP:   ip,
		Port: uPNPPort,
	}
	dialer := net.Dialer{
		Timeout: time.Second * 2,
	}
	conn, err := dialer.Dial("tcp", addr.String())
	if err != nil {
		return nil, err
	}
	// create a line scanner for reading the connection
	// begin scanning at <ESATMUpdate>
	// stop scanning and then close the connection at </ESATMUpdate>
	var data bytes.Buffer
	var scanning bool
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		tmp := scanner.Text()
		if strings.Index(tmp, nasdXMLOpen) != -1 && !scanning {
			scanning = true
		}
		if scanning {
			if n, err := data.WriteString(tmp); err != nil || n != len(tmp) {
				return nil, err
			}
		}
		if strings.Index(tmp, nasdXMLClose) != -1 && scanning {
			break
		}
	}
	conn.Close()
	out = data.Bytes()
	return
}
