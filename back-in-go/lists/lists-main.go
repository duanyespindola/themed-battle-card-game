package lists

import (
	"github.com/duanyespindola/themed-battle-card-game/back-in-go/commons"
	"github.com/duanyespindola/themed-battle-card-game/back-in-go/room"
)

type Lists struct {
	listById                   TListById
	listWaitingForPlayers      TListById
	listWaitingForMatch        TListById
	listWaitingForPlayersOrder []commons.TUUID
}

func (l *Lists) AddToLists(r room.IRoom) {
	if r.RoomStatus() == room.StatusEmpty {
		return
	}
	if r.RoomStatus() == room.StatusWaitingPlayer {
		l.listWaitingForPlayers[r.Id()] = r
		l.listWaitingForPlayersOrder = append(l.listWaitingForPlayersOrder, r.Id())
	}
	if r.RoomStatus() == room.StatusWaitingMatch {
		l.listWaitingForMatch[r.Id()] = r
	}
	l.listById[r.Id()] = r
	for _, player := range r.Players() {
		l.listById[player.Id()] = r
	}
}

func (l *Lists) RemoveFromAllLists(r room.IRoom) {
	delete(l.listById, r.Id())
	delete(l.listWaitingForPlayers, r.Id())
	delete(l.listWaitingForMatch, r.Id())
	for i, id := range l.listWaitingForPlayersOrder {
		if id == r.Id() {
			l.listWaitingForPlayersOrder = append(l.listWaitingForPlayersOrder[:i], l.listWaitingForPlayersOrder[i+1:]...)
			break
		}
	}
	for _, player := range r.Players() {
		delete(l.listById, player.Id())
	}
}

func (l *Lists) FindById(id commons.TUUID) room.IRoom {
	return l.listById[id]
}

func (l *Lists) getOldest(list TListById) room.IRoom {
	for _, id := range l.listWaitingForPlayersOrder {
		if room, exists := list[id]; exists {
			return room
		}
	}
	return nil
}

func (l *Lists) IsThereAnyRoomWaitingForPlayer() room.IRoom {
	if len(l.listWaitingForPlayers) < 1 {
		return nil
	}
	return l.getOldest(l.listWaitingForPlayers)
}

type ListMaker struct{}

func (*ListMaker) NewList() *Lists {
	return &Lists{
		listById:                   make(TListById),
		listWaitingForPlayers:      make(TListById),
		listWaitingForMatch:        make(TListById),
		listWaitingForPlayersOrder: []commons.TUUID{},
	}
}
