package hello_server

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	utils_net "github.com/duanyespindola/themed-battle-card-game/internal/utils"
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

	//find the next available port
	port, err := utils_net.FindNextOpenPort(8080)
	if err != nil {
		t.Fatal(err)
	}

	//start the server
	addr := fmt.Sprintf(":%d", port)
	server := NewServer( addr[1:] ) //remove the ":" from the address
	
	
	if server.Addrs[0] != addr {
		t.Errorf("server address is wrong: got %v want %v", server.Addrs[0], addr)
	}

	if server.Handler == nil {
		t.Error("server handler is nil")
	}
}
