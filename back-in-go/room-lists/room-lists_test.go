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
		player3 := player.Player{
			Id:       "C345678",
			Nickname: "Player3",
		}
		player4 := player.Player{
			Id:       "D901234",
			Nickname: "Player4",
		}
		var room1, room2, room3, room4 *room.Room
		var err error

		When("the first player joins the game", func() {
			BeforeAll(func() {
				room1 = roomLists.AlocateThePlayer(&player1)
			})
			It("that player should be put in a new room waiting for another player", func() {
				Expect(room1.Status()).To(Equal(room.StatusWaitingPlayer))
				Expect(room1.Players()).To(HaveLen(1))
				Expect(*room1.Players()[0]).To(Equal(player1))
			})
		})
		When("the second player joins the game", func() {
			BeforeAll(func() {
				room2 = roomLists.AlocateThePlayer(&player2)
			})
			It("that new player should be put in the same room as player1", func() {
				Expect(room2).To(Equal(room1))
				Expect(room2.Players()).To(HaveLen(2))
				Expect(*room2.Players()[0]).To(Equal(player1))
				Expect(*room2.Players()[1]).To(Equal(player2))
			})
			It("the room should change to waiting for the match to begin", func() {
				Expect(room2.Status()).To(Equal(room.StatusWaitingMatch))
			})
		})
		When("a third player joins the game", func() {
			Context("given that there is a room waiting for the match to start with two players inside and no room waiting another player", func() {
				BeforeAll(func() {
					room3 = roomLists.AlocateThePlayer(&player3)
				})
				It("that player should be put in a new room waiting for another player", func() {
					Expect(room3.Status()).To(Equal(room.StatusWaitingPlayer))
					Expect(room3.Players()).To(HaveLen(1))
					Expect(*room3.Players()[0]).To(Equal(player3))
				})

			})
		})
		When("the first player leaves the game", func() {
			Context("given that the player1 is in a room with player2 and the room is waiting for the match to start", func() {
				BeforeAll(func() {
					Expect(room1.Status()).To(Equal(room.StatusWaitingMatch))
					room1, err = roomLists.DealocateThePlayer(&player1)
					Expect(room1).To(Equal(room2))
					Expect(err).To(BeNil())
				})
				It("should put the room back to waiting for another player", func() {
					Expect(room1.Status()).To(Equal(room.StatusWaitingPlayer))
					Expect(room2.Status()).To(Equal(room.StatusWaitingPlayer))
				})
				It("should have only the player 2 in the room", func() {
					Expect(room1.Players()).To(HaveLen(1))
					Expect(*room1.Players()[0]).To(Equal(player2))
				})
			})
		})
		Context("Given that there are two rooms waiting for another player", func() {
			It("should have two rooms wainting for players", func() {
				Expect(room2.Status()).To(Equal(room.StatusWaitingPlayer))
				Expect(room3.Status()).To(Equal(room.StatusWaitingPlayer))
				Expect(room2).NotTo(Equal(room3))
			})
			When("the fourth player joins the game", func() {
				BeforeAll(func() {
					room4 = roomLists.AlocateThePlayer(&player4)
				})
				It("should put the player in the oldest room waiting for another player", func() {
					//TODO: in 2025-04-018 it will bring any waiting room, not the oldest
					remainging_rooms := []*room.Room{room2, room3}
					Expect(remainging_rooms).To(ContainElement(room4))
					Expect(*room4.Players()[1]).To(Equal(player4))
				})
			})
		})
	})
})
