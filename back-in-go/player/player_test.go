package player_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/duanyespindola/themed-battle-card-game/back-in-go/player"
)

var _ = Describe("Player Unit Test", func() {
	var player1 *player.Player
	var player2 *player.Player

	BeforeEach(func() {
		player1 = player.NewPlayer()
		player2 = player.NewPlayer()
	})

	Context("Contructor", func() {
		It("should create a new player with a random nickname", func() {
			Expect(player1.Nickname).ToNot(BeEmpty())
			Expect(player2.Nickname).ToNot(BeEmpty())
		})
	})

})
