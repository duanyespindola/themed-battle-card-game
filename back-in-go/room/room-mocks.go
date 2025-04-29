package room

import (
	"github.com/duanyespindola/themed-battle-card-game/back-in-go/commons"
	"github.com/duanyespindola/themed-battle-card-game/back-in-go/player"
)

type MockRoomMaker struct{}

func (*MockRoomMaker) NewRoom(id commons.TUUID, status TRoomStatus, players player.TPlayers) (IRoom, error) {
	return &Room{
		id:      id,     //RN-E7DA
		status:  status, //RN-98CB
		players: players,
	}, nil
}
