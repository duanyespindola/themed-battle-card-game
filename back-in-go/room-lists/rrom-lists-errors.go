package roomLists

import "fmt"

var ErrorRoomNotFound = fmt.Errorf("the room was not found in the list")
var ErrorPlayerNotInRoom = fmt.Errorf("the player is not in the room")
