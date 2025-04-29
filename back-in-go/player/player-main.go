package player

import (
	"strings"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/duanyespindola/themed-battle-card-game/back-in-go/commons"
)

/**
* The concrete implementation of the
* IPlayer
**/
type Player struct {
	id       commons.TUUID
	nickname TPlayerNickname
}

func (p *Player) Id() commons.TUUID {
	return p.id
}

func (p *Player) Nickname() TPlayerNickname {
	return p.nickname
}

/**
* Constructor/factory
**/
type PlayerMaker struct{}

func (pm *PlayerMaker) NewPlayer(id commons.TUUID, nickname TPlayerNickname) IPlayer {
	var pm_id commons.TUUID
	if string(id) == "" {
		pm_id = commons.NewUUID()
	} else {
		pm_id = id
	}

	var pm_nickname TPlayerNickname
	if string(nickname) == "" {
		pm_nickname = generateFakeNickname()
	} else {
		pm_nickname = nickname
	}

	return &Player{
		id:       pm_id,
		nickname: pm_nickname,
	}
}

/**
 * Helper functions
 */
func generateFakeNickname() TPlayerNickname {
	nickname := strings.ToLower(gofakeit.Adjective() + "_" + gofakeit.Noun())
	return TPlayerNickname(nickname)
}
