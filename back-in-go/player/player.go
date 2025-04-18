package player

import (
	"strings"

	"github.com/brianvoe/gofakeit/v7"
)

type Player struct {
	Id       TPlayerId
	Nickname TPlayerNickname
}

func NewPlayer() *Player {
	return &Player{
		Id:       TPlayerId(gofakeit.UUID()),
		Nickname: generateFakeNickname(),
	}
}

func generateFakeNickname() TPlayerNickname {
	nickname := strings.ToLower(gofakeit.Adjective() + "_" + gofakeit.Noun())
	return TPlayerNickname(nickname)
}
