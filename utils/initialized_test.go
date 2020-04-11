package utils

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

type testOneInitialized struct {
	initialized bool
}

func (ti testOneInitialized) Initialized() bool {
	return ti.initialized
}

type testTwoInitialized struct {
	initialized bool
}

func (ti testTwoInitialized) Initialized() bool {
	return ti.initialized
}

func TestInitialized(t *testing.T) {

	Convey("Initialized", t, func() {

		Convey("should ignore nil elements", func() {

			So(Initialized(nil, nil), ShouldBeNil)

		})

		Convey("should return error if one of the elements is not initialized", func() {

			err := Initialized(testOneInitialized{initialized: true}, testTwoInitialized{initialized: false})
			So(err, ShouldBeError, "element 'testTwoInitialized' has not been initialized")

		})

		Convey("should return error if one of the pointer elements is not initialized", func() {

			err := Initialized(&testOneInitialized{initialized: true}, &testTwoInitialized{initialized: false})
			So(err, ShouldBeError, "element 'testTwoInitialized' has not been initialized")

		})

		Convey("should no error if all elements are initialized", func() {

			err := Initialized(&testOneInitialized{initialized: true}, &testTwoInitialized{initialized: true})
			So(err, ShouldBeNil)

		})

		Convey("should ignore nil of specific type", func() {

			var s *testOneInitialized = nil

			err := Initialized(s)
			So(err, ShouldBeNil)

		})

	})

}
