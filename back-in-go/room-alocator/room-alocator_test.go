package room_alocator_test

import (
	"github.com/duanyespindola/themed-battle-card-game/back-in-go/lists"
	"github.com/duanyespindola/themed-battle-card-game/back-in-go/player"
	"github.com/duanyespindola/themed-battle-card-game/back-in-go/room"
	room_alocator "github.com/duanyespindola/themed-battle-card-game/back-in-go/room-alocator"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

// factories
var roomAlocatorMaker = &room_alocator.RoomAlocatorMaker{}
var listMaker = &lists.ListMaker{}
var playerMaker = &player.PlayerMaker{}

var roomAlocator = roomAlocatorMaker.NewRoomAlocator(nil)
var list = listMaker.NewList()

var _ = Describe("Room Alocator Unit Test", Ordered, func() {
	Context("AlocatePlayer", func() {
		When("a player joins the game for the first time", func() {
			var player1 player.IPlayer
			BeforeEach(func() {
				player1 = playerMaker.NewPlayer("", "")
			})
			Context("and there is no room waiting for players", func() {
				BeforeEach(func() {
					Expect(list.IsThereAnyRoomWaitingForPlayer()).To(BeNil())
				})
				It("should return a new room, waiting for players, with the player inside", func() {
					r := roomAlocator.AlocatePlayer(player1)

					Expect(r).NotTo(BeNil())
					Expect(r.Players()).To(HaveLen(1))
					Expect(r.Players()).To(ContainElement(player1))
					Expect(r.RoomStatus()).To(Equal(room.StatusWaitingPlayer))
				})
			})
		})
	})
})
