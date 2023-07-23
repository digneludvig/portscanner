package main

import (
	"portscanner/scanner"
)

func main() {
	hostnames := []string{"google.com", "facebook.com", "mail.digne.se"}
	ports := []int{80, 443}
	scanner.ScanRange(hostnames, ports)
}
