package player

import (
	"strings"

	"github.com/brianvoe/gofakeit/v7"
)

type Player struct {
	Nickname string;
	Id string
}

func NewPlayer() *Player {
	return &Player{
		Id: gofakeit.UUID(),
		Nickname: generateFakeNickname(),
	}
}

func generateFakeNickname() string {
	return strings.ToLower(gofakeit.Adjective() + "_" + gofakeit.Noun())
}
