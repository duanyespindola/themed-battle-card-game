package roomManager_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/duanyespindola/themed-battle-card-game/back-in-go/player"
	roomManager "github.com/duanyespindola/themed-battle-card-game/back-in-go/room-manager"
	waitingRoom "github.com/duanyespindola/themed-battle-card-game/back-in-go/waiting-room"
)

var _ = Describe("RoomManager Unit Test", Ordered, func() {
	Context("AlocateThePlayer", func() {
		player1 := player.Player{
			Id:       "A123456",
			Nickname: "Player1",
		}
		// player2 := player.Player{
		// 		Id:       "B789012",
		// 		Nickname: "Player2",
		// }
		// player3 := player.Player{
		// 		Id:       "C345678",
		// 		Nickname: "Player3",
		// }
		var room1 *waitingRoom.WaitingRoom

		When("Asked to find a room for the first player,", func() {

			It("Should return a room waiting for another player", func() {
				room1 = roomManager.AlocateThePlayer(&player1)
				Expect(room1.Status()).To(Equal(waitingRoom.WaitingForAnotherPlayerStatus))
				Expect(room1.Players()).To(HaveLen(1))
				Expect(*room1.Players()[0]).To(Equal(player1))
			})
		})
	})
})
