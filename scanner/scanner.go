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
		fmt.Printf("Port %d is closed or unreachable.\n", port)
		return false
	}

	defer conn.Close()

	fmt.Printf("Port %d is open.\n", port)
	return true
}
