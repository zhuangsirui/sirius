package main

import (
	"github.com/Sirupsen/logrus"
	"github.com/zhuangsirui/sirius/http"
	"github.com/zhuangsirui/sirius/webdav"
)

type simpleAuth struct {
}

func (a simpleAuth) Auth(_, _ string) bool {
	return true
}

func main() {
	server := http.NewServer(http.Config{
		IP:   "127.0.0.1",
		Port: 8000,
	})
	dav := webdav.NewWebDav(webdav.Config{
		Prefix:        "/dav",
		RootDir:       "/tmp/webdav",
		Authenticator: simpleAuth{},
	})
	if err := dav.Init(); err != nil {
		panic(err)
	}
	server.AddHandlerFunc("/dav/", dav.HandlerFunc)
	logrus.WithFields(logrus.Fields{
		"http": "stopped",
	}).Info(server.Serve())
}
