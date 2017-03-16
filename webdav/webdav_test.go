package webdav

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

type authenticator struct {
}

func (a authenticator) Auth(_, _ string) bool {
	return true
}

func TestWebDav(t *testing.T) {
	Convey("webdav", t, func() {

		Convey("init without authenticate", func() {
			wd := NewWebDav(Config{
				Prefix:  "/dav",
				RootDir: "/tmp/webdav",
			})
			err := wd.Init()
			So(err, ShouldNotBeNil)
		})

		Convey("init", func() {
			wd := NewWebDav(Config{
				Prefix:        "/dav",
				RootDir:       "/tmp/webdav",
				Authenticator: authenticator{},
			})
			err := wd.Init()
			So(err, ShouldBeNil)
		})
	})
}