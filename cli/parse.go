package cli

import (
	"flag"
	"fmt"
	"strconv"
	"strings"
)

type Config struct {
	Port int
}

func ParseFlags() *Config {
	portPtr := flag.Int("p", 80, "Port to scan")
	flag.Parse()

	return &Config{
		Port: *portPtr,
	}
}

func parsePorts(input string) []int {
	var ports []int

	portArgs := strings.Split(input, ",")
	for _, portArg := range portArgs {
		if strings.Contains(portArg, "-") {
			rangeStartEnd := strings.Split(portArg, "-")
			if len(rangeStartEnd) == 2 {
				start, err1 := strconv.Atoi(rangeStartEnd[0])
				end, err2 := strconv.Atoi(rangeStartEnd[1])

				if err1 == nil && err2 == nil && start <= end {
					for i := start; i <= end; i++ {
						ports = append(ports, i)
					}
				} else {
					fmt.Printf("Invalid range: %s", portArg)
				}
			} else {
				fmt.Printf("Invalid range: %s", portArg)
			}
		} else {
			port, err := strconv.Atoi(portArg)
			if err != nil {
				fmt.Printf("Invalid port: %s", portArg)
			} else {
				ports = append(ports, port)
			}
		}
	}
	return ports
}
