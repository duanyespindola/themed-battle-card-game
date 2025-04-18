package room

import "github.com/duanyespindola/themed-battle-card-game/back-in-go/player"

// REQ-E973
type TRoomId string

// REQ-B85E
type TRoomStatus int

const (
	StatusWaitingPlayer TRoomStatus = iota + 1
	StatusWaitingMatch
)

type TRoomPlayers []*player.Player
