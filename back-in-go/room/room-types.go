package room

import (
	"github.com/duanyespindola/themed-battle-card-game/back-in-go/commons"
	"github.com/duanyespindola/themed-battle-card-game/back-in-go/player"
)

type TRoomStatus int

const (
	StatusWaitingPlayer TRoomStatus = iota + 1
	StatusWaitingMatch
	StatusEmpty
)

type IWithRoomStatus interface {
	RoomStatus() TRoomStatus
}
type IRoom interface {
	commons.IWithId
	IWithRoomStatus
	player.IWithPlayers
	AddPlayer(p player.IPlayer) error
	RemovePlayer(p player.IPlayer) error
	GetTheOtherPlayer(p player.IPlayer) (player.IPlayer, error)
}

type IRoomMaker interface {
	NewRoom(id commons.TUUID, status TRoomStatus, players player.TPlayers) (IRoom, error)
}
