package roomLists

import "github.com/duanyespindola/themed-battle-card-game/back-in-go/room"

type TListByRoomId map[room.TRoomId]*room.Room

type TListByStatus map[room.TRoomStatus]TListByRoomId
