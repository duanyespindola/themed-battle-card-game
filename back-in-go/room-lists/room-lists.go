package roomLists

/**
 * This class will maintain lists of room bases in
 * a series of indexes to make find a room easier
 *
 * see ./room-lists-mental-model.md to better
 * understand what we want to accomplish
 */

import (
	"github.com/duanyespindola/themed-battle-card-game/back-in-go/player"
	"github.com/duanyespindola/themed-battle-card-game/back-in-go/room"
)

var byStatus = TListByStatus{
	room.StatusWaitingPlayer: make(TListByRoomId),
	room.StatusWaitingMatch:  make(TListByRoomId),
}
var byPlayerId = make(TListByPlayerId)

func AlocateThePlayer(p *player.Player) *room.Room {
	waiting_room := isThereAnyWaitingForPlayer()
	if waiting_room != nil {
		removeFromAllLists(waiting_room)
		_ = waiting_room.AddPlayer(p)
	} else {
		waiting_room = room.NewRoom(p)
	}
	insertOnLists(waiting_room)
	return waiting_room
}

func DealocateThePlayer(p *player.Player) (*room.Room, error) {
	waiting_room, ok := byPlayerId[p.Id]
	if !ok {
		return nil, ErrorRoomNotFound
	}

	_, err := waiting_room.RemovePlayer(p)
	if err != nil {
		return waiting_room, ErrorPlayerNotInRoom
	}

	return waiting_room, nil
}

func isThereAnyWaitingForPlayer() *room.Room {
	waitingRooms := byStatus[room.StatusWaitingPlayer]
	if len(waitingRooms) > 0 {
		wainting_room := getOne(waitingRooms)
		return wainting_room
	}
	return nil
}

func removeFromAllLists(r *room.Room) {
	//remove from byStatus
	delete(byStatus[r.Status()], r.Id())

	//remove from byPlayerId
	for _, p := range r.Players() {
		delete(byPlayerId, p.Id)
	}
}

func insertOnLists(r *room.Room) {
	//update list byStatus
	byStatus[r.Status()][r.Id()] = r

	//update list byPlayerId
	for _, p := range r.Players() {
		byPlayerId[p.Id] = r
	}
}

/**
* maps are unordered, which means we cannot rely on any particular order of elements
* if later we see that a given room is waiting to much and we decide the need of a consistent
* order, whe should consider using a slice to store keys in order or use a struct that
* keeps track of insertion order manually.
* But, for now, any room in the same state do the trick
**/
func getOne(m TListByRoomId) *room.Room {
	for _, value := range m {
		return value
	}
	return nil
}
