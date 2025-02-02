package server

import (
	"fmt"
	"net/http"
)

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello world")
}

func NewServer() *http.Server {
	mux := http.NewServeMux()
	mux.HandleFunc("/", helloWorldHandler)

	return &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
}