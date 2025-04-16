package waitingRoom_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/duanyespindola/themed-battle-card-game/back-in-go/player"
	waitingRoom "github.com/duanyespindola/themed-battle-card-game/back-in-go/waiting-room"
)

/**
* Please, note that this is an "ORDERED" test
* so, be careful when changing the test because
* prior tests will affect all the next ones
**/
var _ = Describe("WaitingRoom Unit Test", Ordered, func() {
	var player1 = player.Player{
		Id:       "521363d-f549-4be8-aaa4-c9fb9fcffb31",
		Nickname: "Player1",
	}
	var player2 = player.Player{
		Id:       "793ad335-2769-4f43-ad13-758a3f2271bc",
		Nickname: "Player2",
	}
	var room waitingRoom.WaitingRoom

	Context("NewWaitingRoom", func() {
		When("Creating a new waiting room", func() {
			BeforeAll(func() {
				room = *waitingRoom.NewWaitingRoom(&player1)
				GinkgoWriter.Printf("Room structure: %+v\n", room)
			})
			It("should has an UID", func() {
				Expect(room.Id()).To(HaveLen(36))
			})
			It("should be in 'waiting-another-player' status", func() {
				Expect(room.Status()).To(Equal(waitingRoom.WaitingForAnotherPlayerStatus))
			})
			It("should have just the player", func() {
				Expect(room.Players()).To(HaveLen(1))
				// Expect(room.Players()[0].Id).To(Equal(player1.Id))
				Expect(*room.Players()[0]).To(Equal(player1))
			})
		})
	})

	Context("AddPlayer", func() {
		BeforeAll(func() {
			room = *waitingRoom.NewWaitingRoom(&player1)
		})

		When("Trying to add the same player again", func() {
			It("should return an error when adding the same player", func() {
				err := room.AddPlayer(&player1)
				Expect(err).To(MatchError("player already in the room"))
			})
		})
		When("Adding a different player", func() {
			It("should allow adding a second player", func() {
				Expect(room.Status()).To(Equal(waitingRoom.WaitingForAnotherPlayerStatus))
				err := room.AddPlayer(&player2)
				Expect(err).To(BeNil())
				Expect(room.Players()).To(HaveLen(2))
			})
			It("should have changed the status to Waiting Match to Starts", func() {
				Expect(room.Status()).To(Equal(waitingRoom.WaitingForMatchToStartStatus))
			})
		})
		When("The status is different from Waiting For Another Player", func() {
			It("should not allow adding new players", func() {
				Expect(room.Status()).NotTo(Equal(waitingRoom.WaitingForAnotherPlayerStatus))
				player3 := player.Player{
					Id:       "B94ag335-2466-3f43-ad13-758a3f2271ee",
					Nickname: "Player3",
				}
				err := room.AddPlayer(&player3)
				Expect(err).To(MatchError("room status does not allow adding players"))

			})

		})

	})

	Context("RemovePlayer", func() {
		When("A room has two players and one player is removed", func() {
			var removedPlayer *player.Player
			var err error
			BeforeAll(func() {
				removedPlayer, err = room.RemovePlayer(&player1)
			})

			It("should put the status back to Waiting For Another Player", func() {
				Expect(err).To(BeNil())
				Expect(room.Status()).To(Equal(waitingRoom.WaitingForAnotherPlayerStatus))
			})
			It("should return the removed player", func() {
				Expect(removedPlayer.Id).To(Equal(player1.Id))
			})
			It("should have just one player in the room", func() {
				Expect(room.Players()).To(HaveLen(1))
			})
		})
	})

})
