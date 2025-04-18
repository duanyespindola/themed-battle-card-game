package roomLists

import (
	"github.com/duanyespindola/themed-battle-card-game/back-in-go/player"
	"github.com/duanyespindola/themed-battle-card-game/back-in-go/room"
)

type TListByRoomId map[room.TRoomId]*room.Room

type TListByStatus map[room.TRoomStatus]TListByRoomId

type TListByPlayerId map[player.TPlayerId]*room.Room
