package main

import (
	"portscanner/cli"
	"portscanner/scanner"
)

func main() {
	config := cli.ParseFlags()
	hostnames := []string{"google.com", "facebook.com", "mail.digne.se"}
	ports := config.Ports
	scanner.ScanRange(hostnames, ports)
}
