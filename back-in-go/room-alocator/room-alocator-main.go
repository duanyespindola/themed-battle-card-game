package room_alocator

import (
	"github.com/duanyespindola/themed-battle-card-game/back-in-go/player"
	"github.com/duanyespindola/themed-battle-card-game/back-in-go/room"
)

type RoomAlocator struct {
	roomMaker room.IRoomMaker
}

func (ra *RoomAlocator) AlocatePlayer(the_player player.IPlayer) room.IRoom {
	r, _ := ra.roomMaker.NewRoom("", room.StatusWaitingPlayer, player.TPlayers{the_player})
	return r
}

/**
 * Factory
 */

type RoomAlocatorMaker struct{}

func (ram *RoomAlocatorMaker) NewRoomAlocator(roomMaker room.IRoomMaker) *RoomAlocator {
	var rm room.IRoomMaker
	if roomMaker == nil {
		rm = &room.RoomMaker{}
	} else {
		rm = roomMaker
	}

	return &RoomAlocator{
		roomMaker: rm,
	}
}
