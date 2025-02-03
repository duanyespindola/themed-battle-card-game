package utils_net

import (
	"fmt"
	"net"
)

func FindNextOpenPort(start int) (int, error) {
	for port := start; port <= 65535; port++ {
		ln, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
		if err == nil {
			ln.Close() // Close immediately to release the port
			return port, nil
		}
	}
	return 0, fmt.Errorf("no open ports found")
}