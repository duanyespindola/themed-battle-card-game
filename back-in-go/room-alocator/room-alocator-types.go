package room_alocator

type IRoomAlocatorMaker interface {
	NewRoomAlocator() *RoomAlocator
}
