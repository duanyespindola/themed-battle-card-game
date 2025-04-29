package room_test

import (
	"github.com/duanyespindola/themed-battle-card-game/back-in-go/commons"
	"github.com/duanyespindola/themed-battle-card-game/back-in-go/player"
	"github.com/duanyespindola/themed-battle-card-game/back-in-go/room"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var roomMaker = &room.RoomMaker{}
var playerMaker = &player.PlayerMaker{}

var _ = Describe("RoomMaker Unit Test", func() {
	Context("NewRoom", func() {
		It("should create an empty room successfully", func() {
			r, err := roomMaker.NewRoom("", room.StatusEmpty, []player.IPlayer{})

			Expect(err).To(BeNil())
			Expect(r.Id()).To(HaveLen(36)) // UUID length
			Expect(r.RoomStatus()).To(Equal(room.StatusEmpty))
			Expect(r.Players()).To(BeEmpty())
		})
		It("should return error when creating a empty room with wrong status", func() {
			r, err := roomMaker.NewRoom("", room.StatusWaitingPlayer, []player.IPlayer{})

			Expect(err).To(MatchError("status must be empty when there are no players"))
			Expect(r).To(BeNil())
		})

		It("should create a room with one player", func() {
			testPlayer := playerMaker.NewPlayer("", "")
			r, err := roomMaker.NewRoom("", room.StatusWaitingPlayer, []player.IPlayer{testPlayer})

			Expect(err).To(BeNil())
			Expect(r.Id()).To(HaveLen(36)) // UUID length
			Expect(r.RoomStatus()).To(Equal(room.StatusWaitingPlayer))
			Expect(r.Players()).To(HaveLen(1))
			Expect(r.Players()[0].Id()).To(Equal(testPlayer.Id()))
		})
		It("should return error when creating a room with one player and wrong status", func() {
			testPlayer := playerMaker.NewPlayer("", "")
			r, err := roomMaker.NewRoom("", room.StatusWaitingMatch, []player.IPlayer{testPlayer})

			Expect(err).To(MatchError("status must be 'waiting for players' when there is only one player"))
			Expect(r).To(BeNil())
		})

		It("should create a room with default values when no parameters are provided", func() {
			testPlayer := playerMaker.NewPlayer("", "")
			r, err := roomMaker.NewRoom("", room.StatusWaitingPlayer, []player.IPlayer{testPlayer})

			Expect(err).To(BeNil())
			Expect(r.Id()).To(HaveLen(36)) // UUID length
			Expect(r.RoomStatus()).To(Equal(room.StatusWaitingPlayer))
			Expect(r.Players()).To(HaveLen(1))
			Expect(r.Players()[0].Id()).To(Equal(testPlayer.Id()))
		})

		It("should create a room with default values when no parameters are provided", func() {
			testPlayer := playerMaker.NewPlayer("", "")
			r, err := roomMaker.NewRoom("", room.StatusWaitingPlayer, []player.IPlayer{testPlayer})

			Expect(err).To(BeNil())
			Expect(r.Id()).To(HaveLen(36)) // UUID length
			Expect(r.RoomStatus()).To(Equal(room.StatusWaitingPlayer))
			Expect(r.Players()).To(HaveLen(1))
			Expect(r.Players()[0].Id()).To(Equal(testPlayer.Id()))
		})
		It("should use provided ID when specified", func() {
			testId := commons.TUUID("test-room-id")
			testPlayer := playerMaker.NewPlayer("", "")
			r, err := roomMaker.NewRoom(testId, room.StatusWaitingPlayer, []player.IPlayer{testPlayer})

			Expect(err).To(BeNil())
			Expect(r.Id()).To(Equal(testId))
		})

		It("should create a room with two players and correct status", func() {
			player1 := playerMaker.NewPlayer("", "")
			player2 := playerMaker.NewPlayer("", "")
			r, err := roomMaker.NewRoom("", room.StatusWaitingMatch, []player.IPlayer{player1, player2})

			Expect(err).To(BeNil())
			Expect(r.Id()).To(HaveLen(36)) // UUID length
			Expect(r.RoomStatus()).To(Equal(room.StatusWaitingMatch))
			Expect(r.Players()).To(HaveLen(2))
			Expect(r.Players()[0].Id()).To(Equal(player1.Id()))
			Expect(r.Players()[1].Id()).To(Equal(player2.Id()))
		})
		It("should return error when creating room with two players but wrong status", func() {
			player1 := playerMaker.NewPlayer("", "")
			player2 := playerMaker.NewPlayer("", "")
			r, err := roomMaker.NewRoom("", room.StatusWaitingPlayer, []player.IPlayer{player1, player2})

			Expect(err).To(MatchError("status must be 'waiting for match' when there are two players"))
			Expect(r).To(BeNil())
		})

		It("should return error when creating room with same player twice", func() {
			player1 := playerMaker.NewPlayer("", "")
			r, err := roomMaker.NewRoom("", room.StatusWaitingPlayer, []player.IPlayer{player1, player1})

			Expect(err).To(MatchError("the same player was informed twice"))
			Expect(r).To(BeNil())
		})

		It("should return error when creating room with more than two players", func() {
			player1 := playerMaker.NewPlayer("", "")
			player2 := playerMaker.NewPlayer("", "")
			player3 := playerMaker.NewPlayer("", "")
			r, err := roomMaker.NewRoom("", room.StatusWaitingMatch, []player.IPlayer{player1, player2, player3})

			Expect(err).To(MatchError("a room can have at most two players"))
			Expect(r).To(BeNil())
		})

	})
})

var _ = Describe("Room Unit Test", func() {
	var testRoom room.IRoom
	var player1, player2, player3 player.IPlayer
	var err error

	BeforeEach(func() {
		player1 = playerMaker.NewPlayer("", "")
		player2 = playerMaker.NewPlayer("", "")
		player3 = playerMaker.NewPlayer("", "")
		testRoom, err = roomMaker.NewRoom("", room.StatusWaitingPlayer, []player.IPlayer{player1})
		Expect(err).To(BeNil())
	})

	Context("AddPlayer", func() {
		It("should allow adding a second player", func() {
			err := testRoom.AddPlayer(player2)
			Expect(err).To(BeNil())
			Expect(testRoom.Players()).To(HaveLen(2))
			Expect(testRoom.RoomStatus()).To(Equal(room.StatusWaitingMatch))
		})
		It("should not allow adding a player if the room is full", func() {
			mockRoomMaker := &room.MockRoomMaker{}
			mockRoom, err := mockRoomMaker.NewRoom(
				commons.TUUID("123"),
				room.StatusEmpty,
				player.TPlayers{})
			err = mockRoom.AddPlayer(player3)
			Expect(err).To(Equal(room.ErrorNotAllowedAddingPlayer))
		})

		It("should not allow adding a player if the room is full", func() {
			mockRoomMaker := &room.MockRoomMaker{}
			mockRoom, err := mockRoomMaker.NewRoom(
				commons.TUUID("123"),
				room.StatusWaitingPlayer,
				player.TPlayers{player1, player2})
			err = mockRoom.AddPlayer(player3)
			Expect(err).To(Equal(room.ErrorNotAllowedAddingPlayer))
		})

		It("should not allow adding the same player twice", func() {
			err := testRoom.AddPlayer(player1)
			Expect(err).To(BeNil())
		})

	})

	Context("RemovePlayer", func() {
		It("should allow removing a player and update the room status", func() {
			_ = testRoom.AddPlayer(player2)
			err := testRoom.RemovePlayer(player2)
			Expect(err).To(BeNil())
			Expect(testRoom.Players()).To(HaveLen(1))
			Expect(testRoom.RoomStatus()).To(Equal(room.StatusWaitingPlayer))
		})

		It("should return an error if the player is not in the room", func() {
			err := testRoom.RemovePlayer(player2)
			Expect(err).To(Equal(room.ErrorPlayerNotInRoom))
		})
		It("should allow removing the last player and update the room status", func() {
			err := testRoom.RemovePlayer(player1)
			Expect(err).To(BeNil())
			Expect(testRoom.Players()).To(HaveLen(0))
			Expect(testRoom.RoomStatus()).To(Equal(room.StatusEmpty))
		})

	})

	Context("GetTheOtherPlayer", func() {
		It("should return the other player in the room", func() {
			_ = testRoom.AddPlayer(player2)
			other_player, err := testRoom.GetTheOtherPlayer(player1)
			Expect(err).To(BeNil())
			Expect(other_player.Id()).To(Equal(player2.Id()))

			other_player, err = testRoom.GetTheOtherPlayer(player2)
			Expect(err).To(BeNil())
			Expect(other_player.Id()).To(Equal(player1.Id()))
		})

		It("should return an error if there is only one player in the room", func() {
			_, err := testRoom.GetTheOtherPlayer(player1)
			Expect(err).To(Equal(room.ErrorJustOnePlayerInRoom))
		})

		It("should return an error if the player is not in the room", func() {
			_ = testRoom.AddPlayer(player2)
			_, err := testRoom.GetTheOtherPlayer(player3)
			Expect(err).To(Equal(room.ErrorPlayerNotInRoom))
		})
	})
})
