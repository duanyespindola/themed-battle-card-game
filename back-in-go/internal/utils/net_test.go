package utils_net

import (
	"net"
	"strconv"
	"testing"
)

func TestFindNextOpenPort(t *testing.T) {
	port, err := FindNextOpenPort(8080)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	ln, err := net.Listen("tcp", ":"+strconv.Itoa(port))
	if err != nil {
		t.Fatalf("Expected to be able to listen on port %d, got %v", port, err)
	}
	ln.Close()
}
