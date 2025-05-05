package room_alocator

import (
	"github.com/duanyespindola/themed-battle-card-game/back-in-go/lists"
	"github.com/duanyespindola/themed-battle-card-game/back-in-go/player"
	"github.com/duanyespindola/themed-battle-card-game/back-in-go/room"
)

type RoomAlocator struct {
	roomMaker room.IRoomMaker
	list      *lists.Lists
}

func (ra *RoomAlocator) AlocatePlayer(the_player player.IPlayer) room.IRoom {
	//The player is already in any room ?
	r := ra.list.FindById(the_player.Id())
	if r != nil {
		return r
	}
	//There is a room waiting for players ??
	r = ra.list.IsThereAnyRoomWaitingForPlayer()
	if r != nil {
		r.AddPlayer(the_player)
	} else {
		r, _ = ra.roomMaker.NewRoom("", room.StatusWaitingPlayer, player.TPlayers{the_player})
	}
	//TODO: transform those two functions into one "UpdateLists" ??
	ra.list.RemoveFromAllLists(r)
	ra.list.AddToLists(r)
	return r
}

func (ra *RoomAlocator) DealocatePlayer(the_player player.IPlayer) room.IRoom {
	r := ra.list.FindById(the_player.Id())
	if r == nil {
		return nil
	}

	ra.list.RemoveFromAllLists(r)
	err := r.RemovePlayer(the_player)
	ra.list.AddToLists(r)

	if err != nil {
		//TODO: send a alert to the developers. It should never occurs this!
	}
	return r
}

/**
 * Factory
 */

type RoomAlocatorMaker struct{}

func (ram *RoomAlocatorMaker) NewRoomAlocator(roomMaker room.IRoomMaker, list *lists.Lists) *RoomAlocator {
	var rm room.IRoomMaker
	if roomMaker == nil {
		rm = &room.RoomMaker{}
	} else {
		rm = roomMaker
	}

	return &RoomAlocator{
		roomMaker: rm,
		list:      list,
	}
}
