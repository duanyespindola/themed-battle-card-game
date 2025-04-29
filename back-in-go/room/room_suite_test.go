package room_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestRoom(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Room Test Suite")
}
