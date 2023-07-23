package scanner

import (
	"fmt"
	"net"
	"time"
)

func ScanPort(hostname string, port int) bool {
	address := fmt.Sprintf("%s:%d", hostname, port)
	conn, err := net.DialTimeout("tcp", address, 5*time.Second)

	if err != nil {
		fmt.Printf("%s\n", err)
		fmt.Printf("\tPort %d is closed or unreachable on %s.\n", port, hostname)
		return false
	}

	defer conn.Close()

	fmt.Printf("\tPort %d is open on %s.\n", port, hostname)
	return true
}

func ScanRange(hostnames []string, ports []int) {
	for _, hostname := range hostnames {
		fmt.Printf("Starting scan on %s, ports: %d\n", hostname, ports)
		for _, port := range ports {
			ScanPort(hostname, port)
		}
		fmt.Printf("\n")
	}
}
