package main

import (
	"portscanner/cli"
	"portscanner/scanner"
)

func main() {
	config := cli.ParseFlags()
	hostnames := []string{"google.com", "facebook.com", "mail.digne.se"}

	scanner.ScanRange(hostnames, config.Ports)
}
