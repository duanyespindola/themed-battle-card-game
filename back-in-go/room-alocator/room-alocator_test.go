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
var playerMaker = &player.PlayerMaker{}
var roomMaker = &room.RoomMaker{}
var listMaker = &lists.ListMaker{}
var roomAlocatorMaker = &room_alocator.RoomAlocatorMaker{}

var roomAlocator *room_alocator.RoomAlocator
var list *lists.Lists

var player1, player2 player.IPlayer
var room1, room2 room.IRoom

var _ = Describe("Room Alocator Unit Test", Ordered, func() {
	BeforeEach(func() {
		list = listMaker.NewList()
		roomAlocator = roomAlocatorMaker.NewRoomAlocator(roomMaker, list)
		player1 = nil
		player2 = nil
		room1 = nil
		room2 = nil
	})

	Context("AlocatePlayer", func() {
		When("a player joins the game for the first time", func() {
			BeforeEach(func() {
				player1 = playerMaker.NewPlayer("111", "player1")
			})
			Context("and there is no room waiting for players", func() {
				BeforeEach(func() {
					Expect(list.IsThereAnyRoomWaitingForPlayer()).To(BeNil())
				})
				It("should return a new room, waiting for players, with the player inside", func() {
					room1 := roomAlocator.AlocatePlayer(player1)

					Expect(room1).NotTo(BeNil())
					Expect(room1.Players()).To(HaveLen(1))
					Expect(room1.Players()).To(ContainElement(player1))
					Expect(room1.RoomStatus()).To(Equal(room.StatusWaitingPlayer))
				})
			})
			Context("and there is a room waiting for players", func() {
				BeforeEach(func() {
					player2 = playerMaker.NewPlayer("222", "player2")
					room2 = roomAlocator.AlocatePlayer(player2)
					Expect(list.IsThereAnyRoomWaitingForPlayer()).To(Equal(room2))
				})
				It("should return that room with updated status and both players inside", func() {
					room1 = roomAlocator.AlocatePlayer(player1)
					Expect(room1).To(Equal(room2))
					Expect(room1.RoomStatus()).To(Equal(room.StatusWaitingMatch))
					Expect(room1.Players()).To(ContainElements(player1, player2))
				})

			})
		})
	})

	Context("DealocatePlayer", func() {
		When("removing a player from a room with two players", func() {
			BeforeEach(func() {
				player1 = playerMaker.NewPlayer("111", "player1")
				player2 = playerMaker.NewPlayer("222", "player2")

				room1 = roomAlocator.AlocatePlayer(player1)
				_ = roomAlocator.AlocatePlayer(player2)
				Expect(room1.RoomStatus()).To(Equal(room.StatusWaitingMatch))
				Expect(room1.Players()).To(ContainElements(player.TPlayers{player1, player2}))
			})
			It("should return a new room, waiting for players, with the other player inside", func() {
				room2 = roomAlocator.DealocatePlayer(player1)

				Expect(room2).NotTo(BeNil())
				Expect(room2.Players()).To(HaveLen(1))
				Expect(room2.Players()).To(ContainElement(player2))
				Expect(room2.RoomStatus()).To(Equal(room.StatusWaitingPlayer))
			})

		})

		When("removing the last player from a room", func() {
			BeforeEach(func() {
				player1 = playerMaker.NewPlayer("111", "player1")
				room1 = roomAlocator.AlocatePlayer(player1)
				Expect(room1.RoomStatus()).To(Equal(room.StatusWaitingPlayer))
				Expect(room1.Players()).To(HaveLen(1))
				Expect(room1.Players()).To(ContainElement(player1))
			})

			It("should return an empty room that is removed from lists", func() {
				room2 = roomAlocator.DealocatePlayer(player1)
				Expect(room2).NotTo(BeNil())
				Expect(room2.RoomStatus()).To(Equal(room.StatusEmpty))
				Expect(room2).To(Equal(room1))

				Expect(list.FindById(room1.Id())).To(BeNil())
				Expect(list.FindById(player1.Id())).To(BeNil())
			})
		})

		When("trying to remove a player that isn't in any room", func() {
			BeforeEach(func() {
				player1 = playerMaker.NewPlayer("111", "player1")
				// Player not allocated
			})

			It("should return nil", func() {
				room1 = roomAlocator.DealocatePlayer(player1)
				Expect(room1).To(BeNil())
			})
		})
	})

})
