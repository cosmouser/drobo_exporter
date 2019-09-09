package main

import (
	"bufio"
	"bytes"
	"net"
	"strings"
	"time"
)

const (
	nasdXMLOpen  = "<ESATMUpdate>"
	nasdXMLClose = "</ESATMUpdate>"
)

func dialUPNP(target string) (out []byte, err error) {
	host, port, err := net.SplitHostPort(target)
	if err != nil {
		return nil, err
	}
	dialer := net.Dialer{
		Timeout: time.Second * 2,
	}
	conn, err := dialer.Dial("tcp", net.JoinHostPort(host, port))
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
