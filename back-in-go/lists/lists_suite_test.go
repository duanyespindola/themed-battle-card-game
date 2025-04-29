package lists_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestRoomLists(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Lists Suite")
}
