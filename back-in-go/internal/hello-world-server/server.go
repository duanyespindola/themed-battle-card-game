package server

import (
	"fmt"
	"net"
	"net/http"

	"github.com/lesismal/nbio/nbhttp"
)

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello world")
}

func NewServer(port string) *nbhttp.Server {
	mux := http.NewServeMux()
	mux.HandleFunc("/", helloWorldHandler)

	svr := nbhttp.NewServer(nbhttp.Config{
		Network: "tcp",
		Addrs:  []string{":" + port},
		Handler: mux,
	})
	
	return svr
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