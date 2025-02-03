package server

import (
	"fmt"
	"net"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/joho/godotenv"
)

func TestHelloWorldHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(helloWorldHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := "Hello world"
	if rr.Body.String() != expected {
		t.Errorf("Handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}

func TestNewServer(t *testing.T) {
	//read the .env file
	config, err := godotenv.Read("../../.env")
	if err != nil {
		t.Fatal(err)
	}

	//check if the PORT is setted 
	start_port := config["PORT"]
	if start_port == "" {
		start_port = "8080"
	}

	//convert the PORT to int
	port, err := strconv.Atoi(start_port)
	if err != nil {
		t.Fatal(err)
	}

	//find the next available port
	port, err = FindNextOpenPort(port)
	if err != nil {
		t.Fatal(err)
	}

	//start the server
	server := NewServer( fmt.Sprintf("%d", port) )
	if server.Addr != fmt.Sprintf(":%d", port) {
		t.Errorf("server address is wrong: got %v want %v", server.Addr, ":8080")
	}

	if server.Handler == nil {
		t.Error("server handler is nil")
	}
}
func TestFindNextOpenPort(t *testing.T) {
	startPort := 8080
	port, err := FindNextOpenPort(startPort)
	if err != nil {
		t.Fatalf("Expected to find an open port, but got error: %v", err)
	}

	if port < startPort  {
		t.Errorf("Found port %d is lesse than startPort %d", port, startPort)
	}

	// Try to listen on the found port to ensure it's actually open
	ln, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		t.Fatalf("Expected port %d to be open, but got error: %v", port, err)
	}
	ln.Close()
}
