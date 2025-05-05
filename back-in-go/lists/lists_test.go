package lists

import (
	"github.com/duanyespindola/themed-battle-card-game/back-in-go/player"
	"github.com/duanyespindola/themed-battle-card-game/back-in-go/room"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var roomMaker = &room.MockRoomMaker{}
var playerMaker = &player.PlayerMaker{}
var listMaker = &ListMaker{}

var lists *Lists

// var listById *TListById
// var listWaitingForPlayers *TListById
// var listWaitingForMatch *TListById
// var listWaitingForPlayersOrder *[]commons.TUUID

var _ = Describe("Lists Unit Test", func() {
	player1 := playerMaker.NewPlayer("", "")
	player2 := playerMaker.NewPlayer("", "")

	BeforeEach(func() {
		// always start with empty lists
		lists = listMaker.NewList()
	})

	Describe("AddToLists", func() {

		It("should add a room to the listById", func() {
			mockRoom, _ := roomMaker.NewRoom("room1", room.StatusWaitingMatch, []player.IPlayer{player1, player2})

			lists.AddToLists(mockRoom)

			Expect(lists.listById).To(HaveLen(3))
			Expect(lists.listById).To(HaveKey(mockRoom.Id()))
			Expect(lists.listById).To(HaveKey(player1.Id()))
			Expect(lists.listById).To(HaveKey(player2.Id()))
		})

		It("should add a room to the listWaitingForPlayers", func() {
			mockRoom, _ := roomMaker.NewRoom("room1", room.StatusWaitingPlayer, []player.IPlayer{player1})

			lists.AddToLists(mockRoom)

			Expect(lists.listWaitingForPlayers).To(HaveLen(1))
			Expect(lists.listWaitingForMatch).To(HaveLen(0))
			Expect(lists.IsThereAnyRoomWaitingForPlayer()).To(Equal(mockRoom))
			Expect(lists.listWaitingForPlayers).To(HaveKey(mockRoom.Id()))
		})

		It("should add a room to the listWaitingForMatch", func() {
			mockRoom, _ := roomMaker.NewRoom("room1", room.StatusWaitingMatch, []player.IPlayer{player1, player2})

			lists.AddToLists(mockRoom)

			Expect(lists.listWaitingForMatch).To(HaveLen(1))
			Expect(lists.listWaitingForPlayers).To(HaveLen(0))

			Expect(lists.listWaitingForMatch).To(HaveKey(mockRoom.Id()))
		})

		It("should not add a room with status Empty", func() {
			mockRoom, _ := roomMaker.NewRoom("room1", room.StatusEmpty, []player.IPlayer{})

			lists.AddToLists(mockRoom)

			Expect(lists.listWaitingForPlayers).To(HaveLen(0))
			Expect(lists.listWaitingForMatch).To(HaveLen(0))
			Expect(lists.listWaitingForPlayersOrder).To(HaveLen(0))
		})

	})

	Describe("RemoveFromAllLists", func() {
		It("should remove a room from all lists", func() {
			mockRoom1, _ := roomMaker.NewRoom("room1", room.StatusWaitingPlayer, []player.IPlayer{player1})
			lists.AddToLists(mockRoom1)

			palyer3 := playerMaker.NewPlayer("", "")
			mockRoom2, _ := roomMaker.NewRoom("room1", room.StatusWaitingMatch, []player.IPlayer{player2, palyer3})
			lists.AddToLists(mockRoom2)

			Expect(lists.listById).To(HaveKey(mockRoom1.Id()))
			Expect(lists.listById).To(HaveKey(palyer3.Id()))
			Expect(lists.listWaitingForPlayers).To(HaveKey(mockRoom1.Id()))
			Expect(lists.listWaitingForMatch).To(HaveKey(mockRoom2.Id()))

			lists.RemoveFromAllLists(mockRoom1)

			Expect(lists.listById).NotTo(HaveKey(mockRoom1.Id()))
			Expect(lists.listById).NotTo(HaveKey(player1.Id()))
			Expect(lists.listWaitingForPlayers).To(HaveLen(0))

			lists.RemoveFromAllLists(mockRoom2)

			Expect(lists.listById).NotTo(HaveKey(mockRoom2.Id()))
			Expect(lists.listById).NotTo(HaveKey(palyer3.Id()))
			Expect(lists.listWaitingForMatch).To(HaveLen(0))
		})
	})

	Describe("IsThereAnyRoomWaitingForPlayer", func() {
		It("should return the oldest room in the list", func() {
			Expect(lists.IsThereAnyRoomWaitingForPlayer()).To(BeNil())

			mockRoom1, _ := roomMaker.NewRoom("room1", room.StatusWaitingPlayer, []player.IPlayer{player1})
			lists.AddToLists(mockRoom1)

			mockRoom2, _ := roomMaker.NewRoom("room2", room.StatusWaitingPlayer, []player.IPlayer{player2})
			lists.AddToLists(mockRoom2)

			foundRoom := lists.IsThereAnyRoomWaitingForPlayer()
			Expect(mockRoom1).To(Equal(foundRoom))
		})
		It("should return nil if there is no room in the list", func() {
			oldestRoom := lists.IsThereAnyRoomWaitingForPlayer()
			Expect(oldestRoom).To(BeNil())
		})

	})

	Describe("FindById", func() {
		It("should a room it's id or the players ids", func() {
			mockRoom, _ := roomMaker.NewRoom("room1", room.StatusWaitingMatch, []player.IPlayer{player1, player2})
			lists.AddToLists(mockRoom)

			//Room created but not added
			player3 := playerMaker.NewPlayer("", "")
			room2, _ := roomMaker.NewRoom("", room.StatusWaitingPlayer, []player.IPlayer{player3})

			Expect(lists.listById).To(HaveLen(3))
			Expect(lists.FindById(mockRoom.Id())).To(Equal(mockRoom))
			Expect(lists.FindById(player1.Id())).To(Equal(mockRoom))
			Expect(lists.FindById(player2.Id())).To(Equal(mockRoom))
			Expect(lists.FindById(room2.Id())).To(BeNil())
			Expect(lists.FindById(player3.Id())).To(BeNil())
		})
	})

	Describe("getOldest", func() {
		It("should return the oldest room in the list", func() {
			mockRoom1, _ := roomMaker.NewRoom("room1", room.StatusWaitingPlayer, []player.IPlayer{player1})
			lists.AddToLists(mockRoom1)

			mockRoom2, _ := roomMaker.NewRoom("room2", room.StatusWaitingPlayer, []player.IPlayer{player2})
			lists.AddToLists(mockRoom2)

			oldestRoom := lists.getOldest(lists.listWaitingForPlayers)
			Expect(oldestRoom).To(Equal(mockRoom1))
		})
		It("should return nil if there is no room in the list", func() {
			oldestRoom := lists.getOldest(lists.listWaitingForPlayers)
			Expect(oldestRoom).To(BeNil())
		})

	})

})
