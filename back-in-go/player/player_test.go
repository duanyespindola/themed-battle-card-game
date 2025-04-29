package player_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/duanyespindola/themed-battle-card-game/back-in-go/commons"
	"github.com/duanyespindola/themed-battle-card-game/back-in-go/player"
)

var _ = Describe("Player Unit Test", func() {
	var playerMaker *player.PlayerMaker

	BeforeEach(func() {
		playerMaker = &player.PlayerMaker{}
	})

	Context("When creating a new player", func() {
		It("should generate UUID and nickname when no parameters are provided", func() {
			p := playerMaker.NewPlayer("", "")
			Expect(p.Id()).To(HaveLen(36))
			Expect(p.Nickname()).ToNot(BeEmpty())
		})

		It("should use provided ID and generate nickname when only ID is provided", func() {
			testId := commons.TUUID("test-uuid-123")
			p := playerMaker.NewPlayer(testId, "")
			Expect(p.Id()).To(Equal(testId))
			Expect(p.Nickname()).ToNot(BeEmpty())
		})

		It("should generate UUID and use provided nickname when only nickname is provided", func() {
			testNickname := player.TPlayerNickname("test_player")
			p := playerMaker.NewPlayer("", testNickname)
			Expect(p.Id()).To(HaveLen(36))
			Expect(p.Nickname()).To(Equal(testNickname))
		})

		It("should use both provided ID and nickname when both are provided", func() {
			testId := commons.TUUID("test-uuid-123")
			testNickname := player.TPlayerNickname("test_player")
			p := playerMaker.NewPlayer(testId, testNickname)
			Expect(p.Id()).To(Equal(testId))
			Expect(p.Nickname()).To(Equal(testNickname))
		})
	})
})
