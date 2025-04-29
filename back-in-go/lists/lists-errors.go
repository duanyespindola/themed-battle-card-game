package lists

import "fmt"

var ErrorRoomNotFound = fmt.Errorf("the room was not found")
var ErrorPlayerNotInRoom = fmt.Errorf("the player is not in the room")
