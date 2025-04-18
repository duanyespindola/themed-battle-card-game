package roomLists_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/duanyespindola/themed-battle-card-game/back-in-go/player"
	roomLists "github.com/duanyespindola/themed-battle-card-game/back-in-go/room-lists"
	waitingRoom "github.com/duanyespindola/themed-battle-card-game/back-in-go/waiting-room"
)

var _ = Describe("RoomLists Unit Test", Ordered, func() {
	Context("AlocateThePlayer", func() {
		player1 := player.Player{
			Id:       "A123456",
			Nickname: "Player1",
		}
		player2 := player.Player{
			Id:       "B789012",
			Nickname: "Player2",
		}
		// player3 := player.Player{
		// 		Id:       "C345678",
		// 		Nickname: "Player3",
		// }
		var room1 *waitingRoom.WaitingRoom

		When("Asked to find a room for the first player,", func() {
			room := roomLists.AlocateThePlayer(&player1)
			It("Should return a room with status 'Waiting For Another Player'", func() {
				Expect(room.Status()).To(Equal(waitingRoom.WaitingForAnotherPlayerStatus))
				Expect(room.Players()).To(HaveLen(1))
				Expect(*room.Players()[0]).To(Equal(player1))
			})
			room1 = room
		})
		When("Asked to find a room for the second player,", func() {
			room := roomLists.AlocateThePlayer(&player2)
			It("Should return the same room of the first player with status 'Waiting For Match To Start'", func() {
				Expect(room).To(Equal(room1))
				Expect(room.Status()).To(Equal(waitingRoom.WaitingForMatchToStartStatus))
				Expect(room.Players()).To(HaveLen(2))
				Expect(*room.Players()[0]).To(Equal(player1))
				Expect(*room.Players()[1]).To(Equal(player2))
			})
		})
	})
})
