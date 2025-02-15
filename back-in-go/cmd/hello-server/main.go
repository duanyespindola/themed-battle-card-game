package main

import (
	"log"

	server "github.com/duanyespindola/themed-battle-card-game/internal/hello-server"
	"github.com/joho/godotenv"
)

func main() {
	config, error := godotenv.Read()
	if error != nil {
    	log.Fatal("Error loading .env file")
  	}
	s := server.NewServer(config["PORT"])
	s.Start()
	defer s.Stop()
	s.Wait()
}