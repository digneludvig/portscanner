package main

import (
	"portscanner/scanner"
)

func main() {
	hostname := "google.com"
	port := 80

	scanner.ScanPort(hostname, port)
}
