package main_test

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestSpec2(t *testing.T) {

	// Only pass t into top-level Convey calls
	x := 2
	Convey(fmt.Sprintf("Given the number %d as a starting value", x), t, func() {

		Convey("When the number is incremented", func() {
			x++

			Convey("The value should be greater by one", func() {
				So(x, ShouldEqual, 3)
			})
		})
	})
}