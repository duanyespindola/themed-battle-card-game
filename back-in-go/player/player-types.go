package player

import (
	"github.com/duanyespindola/themed-battle-card-game/back-in-go/commons"
)

type TPlayerNickname string
type IWithNickname interface {
	Nickname() TPlayerNickname
}

type IPlayer interface {
	commons.IWithId
	IWithNickname
}

type TPlayers []IPlayer
type IWithPlayers interface {
	Players() TPlayers
}

type IPlayerMaker interface {
	NewPlayer(id commons.TUUID, nickname TPlayerNickname) IPlayer
}
