package echo_server

import (
	"fmt"
	"log"
	"net/url"
	"testing"
	"time"

	"github.com/lesismal/nbio/nbhttp"
	"github.com/lesismal/nbio/nbhttp/websocket"

	utils_net "github.com/duanyespindola/themed-battle-card-game/internal/utils"
)

var (
	server_engine *nbhttp.Engine
	client_engine *nbhttp.Engine
	connection *websocket.Conn

	host string = "localhost"
	start_port int = 8080
	port int
	path string = "/ws"
	
	test_is_done chan bool = make(chan bool,1)
	ready_to_write chan bool = make(chan bool,1)

	message_to_send string = "Hello, from test"
	want string = "echo: "+message_to_send
	test *testing.T
)

func TestEchoServer(t *testing.T) {
	test = t
	startServer()
	connectClient()

	<-ready_to_write
	connection.WriteMessage(websocket.TextMessage, []byte(message_to_send))

	// this will wait the test to finish
	// in the onMessage callback in the Upgrader
	<-test_is_done
}

func newClientUpgrader() *websocket.Upgrader {
	u := websocket.NewUpgrader()
	u.OnMessage(func(c *websocket.Conn, messageType websocket.MessageType, data []byte) {
		got := string(data)
		test_is_done <- true
		if( got != want) {
			test.Errorf("got %+v, want %+v", got, want)
		}
	})

	return u
}

func startServer() {
	test.Helper()
	//find the next available port
	open_port, err := utils_net.FindNextOpenPort(start_port)
	port = open_port
	if err != nil {
		test.Fatal("FindNextOpenPort :", err)
	}

	srv := NewEchoServer(host, port, path)
	srv.Start()
}

func connectClient() {
	test.Helper()
	host_url := fmt.Sprintf("%s:%d", host, port)
	u := url.URL{Scheme: "ws", Host: host_url, Path: path}
	log.Printf("connecting to %s", u.String())
	client_engine := nbhttp.NewEngine(nbhttp.Config{}) //empty config creates a client
	err := client_engine.Start()
	if err != nil {
		test.Fatal("client_engine.Start: ", err)
		return
	}

	dialer := &websocket.Dialer{
		Engine:      client_engine,
		Upgrader:    newClientUpgrader(),
		DialTimeout: time.Second * 3,
	}
	c, _, err := dialer.Dial(u.String(), nil)
	if err != nil {
		test.Fatal("dialer.Dial :", err)
	}
	connection = c
	ready_to_write <- true
}


