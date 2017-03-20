package main

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestAuthenticator(t *testing.T) {
	Convey("authenticator", t, func() {
		a := new(Authenticator)
		users := []User{
			User{
				Username: "hello",
				Password: "world",
			},
		}
		a.SetUsers(users)
		So(a.Auth("hello", "world"), ShouldBeTrue)
		So(a.Auth("world", "hello"), ShouldBeFalse)
	})
}
