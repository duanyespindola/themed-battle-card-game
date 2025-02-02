package main

import server "github.com/duanyespindola/themed-battle-card-game/internal/hello-world-server"

func main() {
	s := server.NewServer()
	s.ListenAndServe()
	defer s.Close()
}