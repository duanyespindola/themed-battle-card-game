package roomLists_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/duanyespindola/themed-battle-card-game/back-in-go/player"
	"github.com/duanyespindola/themed-battle-card-game/back-in-go/room"
	roomLists "github.com/duanyespindola/themed-battle-card-game/back-in-go/room-lists"
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
		var room1 *room.Room
		When("Asked to find a room for the first player,", func() {
			It("Should return a room with status 'Waiting For Another Player'", func() {
				room1 = roomLists.AlocateThePlayer(&player1)
				Expect(room1.Status()).To(Equal(room.StatusWaitingPlayer))
				Expect(room1.Players()).To(HaveLen(1))
				Expect(*room1.Players()[0]).To(Equal(player1))
			})
		})
		When("Asked to find a room for the second player,", func() {
			It("Should return the same room of the first player with status 'Waiting For Match To Start'", func() {
				room2 := roomLists.AlocateThePlayer(&player2)
				Expect(room2).To(Equal(room1))
				Expect(room2.Status()).To(Equal(room.StatusWaitingMatch))
				Expect(room2.Players()).To(HaveLen(2))
				Expect(*room2.Players()[0]).To(Equal(player1))
				Expect(*room2.Players()[1]).To(Equal(player2))
			})
		})
	})
})
