package room

import (
	"fmt"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/duanyespindola/themed-battle-card-game/back-in-go/player"
)

type Room struct {
	id      TRoomId
	status  TRoomStatus
	players TRoomPlayers
}

func (w *Room) Id() TRoomId {
	return TRoomId(w.id)
}
func (w *Room) Status() TRoomStatus {
	return TRoomStatus(w.status)
}
func (w *Room) Players() TRoomPlayers {
	return TRoomPlayers(w.players)
}
func (w *Room) AddPlayer(p *player.Player) error {
	if w.playerInTheRoom(p) {
		return fmt.Errorf("player already in the room")
	}
	if w.status != StatusWaitingPlayer {
		return fmt.Errorf("room status does not allow adding players")
	}

	w.status = StatusWaitingMatch
	w.players = append(w.players, p)
	return nil
}

func (w *Room) playerInTheRoom(p *player.Player) bool {
	exists := false
	//checando se é o player 1
	if len(w.players) > 0 {
		exists = (w.players[0].Id == p.Id)
	}
	//checando se é o player 2
	if !exists && len(w.players) > 1 {
		exists = (w.players[1].Id == p.Id)
	}
	return exists
}

func NewRoom(p *player.Player) *Room {
	return &Room{
		id:      TRoomId(gofakeit.UUID()),
		status:  StatusWaitingPlayer,
		players: []*player.Player{p},
	}
}

func (room *Room) RemovePlayer(player_to_remove *player.Player) (*player.Player, error) {
	if !room.playerInTheRoom(player_to_remove) {
		return nil, fmt.Errorf("player not found in the room")
	}

	var idx int //player index on the room.players array
	if room.players[0].Id == player_to_remove.Id {
		idx = 0
	} else {
		idx = 1
	}

	removed_player := room.players[idx]
	room.players = append(room.players[:idx], room.players[idx+1:]...)

	room.status = StatusWaitingPlayer

	return removed_player, nil
}
