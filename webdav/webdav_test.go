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
				Prefix:  "/dav/",
				RootDir: "/tmp/webdav",
			})
			err := wd.Init()
			So(err, ShouldNotBeNil)
		})

		Convey("init", func() {
			wd := NewWebDav(Config{
				Prefix:        "/dav/",
				RootDir:       "/tmp/webdav",
				Authenticator: authenticator{},
			})
			err := wd.Init()
			So(err, ShouldBeNil)

			Convey("ensure user dir", func() {
				path, err := wd.ensureUserDir("sirius")
				So(path, ShouldNotEqual, "")
				So(err, ShouldBeNil)
			})
		})

		Convey("ensure handler", func() {
			wd := NewWebDav(Config{
				Prefix:        "/dav/",
				RootDir:       "/tmp/webdav",
				Authenticator: authenticator{},
			})
			handler, err := wd.ensureHandler("sirius")
			So(err, ShouldBeNil)
			So(handler, ShouldNotBeNil)
		})
	})
}
