package server

import (
	"fmt"
	"net"
	"net/http"
)

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello world")
}

func NewServer(port string) *http.Server {
	mux := http.NewServeMux()
	mux.HandleFunc("/", helloWorldHandler)

	return &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}
}

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