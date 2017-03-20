package main

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestConfig(t *testing.T) {
	Convey("config", t, func() {
		c := new(Config)
		err := c.ParseFile("sirius.toml")
		So(err, ShouldBeNil)
	})
}
