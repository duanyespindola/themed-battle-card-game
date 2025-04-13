package player

import (
	"math/rand"
)

type Player struct {
	Nickname string
}

func NewPlayer() *Player {
	return &Player{
		Nickname: generateRandomNickname(),
	}
}

func generateRandomNickname() string {
	adjectives := []string{"Swift", "Brave", "Clever", "Mighty"}
	nouns := []string{"Tiger", "Eagle", "Shark", "Wolf"}
	return adjectives[rand.Intn(len(adjectives))] + nouns[rand.Intn(len(nouns))]
}
