package room

import "fmt"

var ErrorPlayerAlreadyInRoom = fmt.Errorf("player already in the room")
var ErrorPlayerNotInRoom = fmt.Errorf("the player is not in the room")
var ErrorNotAllowedAddingPlayer = fmt.Errorf("room status does not allow adding players")
