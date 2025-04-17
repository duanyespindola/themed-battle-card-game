package roomManager_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestRoomManager(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "RoomManager Suite")
}
