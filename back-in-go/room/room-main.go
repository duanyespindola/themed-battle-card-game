package room

import (
	"fmt"
	"slices"

	"github.com/duanyespindola/themed-battle-card-game/back-in-go/commons"
	"github.com/duanyespindola/themed-battle-card-game/back-in-go/player"
)

type Room struct {
	id      commons.TUUID
	status  TRoomStatus
	players player.TPlayers
}

/**
* Implementing interfaces:
* IWithId
* IWithRoomStatus
* IWithPlayers
* IRoom
**/
func (r *Room) Id() commons.TUUID {
	return r.id
}
func (r *Room) RoomStatus() TRoomStatus {
	return r.status
}
func (r *Room) Players() player.TPlayers {
	return r.players
}

// Other functions
func (r *Room) IsThePlayerInTheRoom(p player.IPlayer) bool {
	return r.getPlayerIndex(p) != -1
}

func (r *Room) AddPlayer(p player.IPlayer) error {
	if r.status != StatusWaitingPlayer {
		return ErrorNotAllowedAddingPlayer
	}
	if len(r.Players()) > 1 {
		return ErrorNotAllowedAddingPlayer
	}
	if r.IsThePlayerInTheRoom(p) {
		return nil
	}

	r.players = append(r.players, p)
	if len(r.Players()) > 1 {
		r.status = StatusWaitingMatch
	}
	return nil
}

func (r *Room) RemovePlayer(p player.IPlayer) error {
	var idx int = r.getPlayerIndex(p)

	if idx < 0 {
		return ErrorPlayerNotInRoom
	}
	r.players = slices.Delete(r.players, idx, idx+1)

	if len(r.players) < 1 {
		r.status = StatusEmpty
	} else {
		r.status = StatusWaitingPlayer
	}

	return nil
}

func (r *Room) GetTheOtherPlayer(p player.IPlayer) (player.IPlayer, error) {
	idx := r.getPlayerIndex(p)
	if idx < 0 {
		return nil, ErrorPlayerNotInRoom
	}
	if len(r.Players()) < 2 {
		return nil, ErrorJustOnePlayerInRoom
	}
	if idx == 0 {
		return r.Players()[1], nil
	} else {
		return r.Players()[0], nil
	}

}

func (r *Room) getPlayerIndex(p player.IPlayer) int {
	idx := -1

	// Check if it's player 1
	if len(r.players) > 0 && r.players[0].Id() == p.Id() {
		idx = 0
	}
	// Check if it's player 2
	if idx < 0 && len(r.players) > 1 && r.players[1].Id() == p.Id() {
		idx = 1
	}
	return idx
}

/**
 * Factory
 * Implementing interface IRoomMaker
 */
type RoomMaker struct{}

func (*RoomMaker) NewRoom(id commons.TUUID, status TRoomStatus, players player.TPlayers) (IRoom, error) {
	if len(players) == 0 && status != StatusEmpty {
		return nil, fmt.Errorf("status must be empty when there are no players")
	}
	if len(players) == 1 && status != StatusWaitingPlayer {
		return nil, fmt.Errorf("status must be 'waiting for players' when there is only one player")
	}
	if len(players) == 2 && players[0].Id() == players[1].Id() {
		return nil, fmt.Errorf("the same player was informed twice")
	}
	if len(players) == 2 && status != StatusWaitingMatch {
		return nil, fmt.Errorf("status must be 'waiting for match' when there are two players")
	}
	if len(players) > 2 {
		return nil, fmt.Errorf("a room can have at most two players")
	}

	var room_id commons.TUUID
	if id == "" {
		room_id = commons.NewUUID()
	} else {
		room_id = id
	}

	return &Room{
		id:      room_id, //RN-E7DA
		status:  status,  //RN-98CB
		players: players,
	}, nil
}
