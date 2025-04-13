package app2_test

import (
	"testing"
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)


var _ = Describe("Unit Test App", func() {
	BeforeEach(func() {
		// fmt.Println("External BeforeEach")
	})

	Describe("Test 1", func() {
		BeforeEach(func() {
			// fmt.Println("Internal BeforeEach")
		})
		It("should run the test1", func() {
			// fmt.Println("Running test1")
			time.Sleep(2 * time.Second)
			Expect(1).To(Equal(21))
		})
		It("should run the test1.2", func() {
			// fmt.Println("Running test1.2")
			Expect(1).To(Equal(1))
		})
	})
	Describe("Test 2", Label("so-eu"),func() {
		It("should run the test2", func() {
			// fmt.Println("Running test1.2")

		})
	})
})


func TestApp(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "App Suite")
}
