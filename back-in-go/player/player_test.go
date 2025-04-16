package player_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/duanyespindola/themed-battle-card-game/back-in-go/player"
)

var _ = Describe("Player Unit Test", func() {
	var player1 *player.Player

	BeforeEach(func() {
		player1 = player.NewPlayer()
	})

	Context("Contructor", func() {
		It("should have an id of type UUID", func() {
			Expect(player1.Id).To(HaveLen(36))
		})
		It("should a random nickname", func() {
			Expect(player1.Nickname).ToNot(BeEmpty())
		})

	})

})
