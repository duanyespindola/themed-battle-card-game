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
	waitingRoom "github.com/duanyespindola/themed-battle-card-game/back-in-go/waiting-room"
)

type RoomList struct {
	byStatus map[waitingRoom.Status][]*waitingRoom.WaitingRoom
}

var roomList = RoomList{
	byStatus: make(map[waitingRoom.Status][]*waitingRoom.WaitingRoom),
}

func thereIsAnyWaitingForAnotherPlayer() *waitingRoom.WaitingRoom {
	waitingRooms := roomList.byStatus[waitingRoom.WaitingForAnotherPlayerStatus]
	if len(waitingRooms) > 0 {
		return waitingRooms[0]
	}
	return nil
}

func AlocateThePlayer(p *player.Player) *waitingRoom.WaitingRoom {
	room := thereIsAnyWaitingForAnotherPlayer()
	if room != nil {
		roomList.byStatus[waitingRoom.WaitingForAnotherPlayerStatus] = roomList.byStatus[waitingRoom.WaitingForAnotherPlayerStatus][1:]
		_ = room.AddPlayer(p)
		return room
	}
	room = waitingRoom.NewWaitingRoom(p)
	roomList.byStatus[waitingRoom.WaitingForAnotherPlayerStatus] = append(roomList.byStatus[waitingRoom.WaitingForAnotherPlayerStatus], room)
	return room
}
