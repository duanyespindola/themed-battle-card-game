package room_alocator_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestPlayer(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Room Alocator Test Suite")
}
