package lists

import (
	"github.com/duanyespindola/themed-battle-card-game/back-in-go/commons"
	"github.com/duanyespindola/themed-battle-card-game/back-in-go/room"
)

// Updated TListById to accept room.IRoom instead of *room.Room
type TListById map[commons.TUUID]room.IRoom

type IListMaker interface {
	NewList() *Lists
}
