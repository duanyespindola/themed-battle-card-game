package waitingRoom_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestWaitingRoom(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "WaitingRoom Test Suite")
}
