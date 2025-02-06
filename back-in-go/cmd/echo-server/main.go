package main

import (
	"context"
	"fmt"
	"time"

	echo_server "github.com/duanyespindola/themed-battle-card-game/internal/echo-server"
)


func main() {
	srv := echo_server.NewEchoServer("localhost", 8080, "/ws")
	
	err := srv.Start()
	if err != nil {
		fmt.Printf("nbio.Start failed: %v\n", err)
		return
	}

	srv.Wait()

	// interrupt := make(chan os.Signal, 1)
	// signal.Notify(interrupt, os.Interrupt)
	// <-interrupt

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()
	srv.Shutdown(ctx)
}