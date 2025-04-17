package roomManager

/**
 * This class will maintain and store
 * the rooms in indexed lists to
 * facilitate quering and retrieval
 */
import (
	"github.com/duanyespindola/themed-battle-card-game/back-in-go/player"
	waitingRoom "github.com/duanyespindola/themed-battle-card-game/back-in-go/waiting-room"
)

func AlocateThePlayer(p *player.Player) *waitingRoom.WaitingRoom {
	newRoom := waitingRoom.NewWaitingRoom(p)
	return newRoom
}
