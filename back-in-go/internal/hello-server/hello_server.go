package hello_server

import (
	"fmt"
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
