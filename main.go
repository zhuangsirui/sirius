package main

import (
	"flag"
	"os"
	"os/signal"
	"sync"

	"github.com/Sirupsen/logrus"
	"github.com/zhuangsirui/sirius/http"
	"github.com/zhuangsirui/sirius/webdav"
)

var (
	config        = new(Config)
	authenticator = new(Authenticator)
)

func init() {
	var c string
	flag.StringVar(&c, "c", "sirius.toml", "sirius -c=/path/to/config.toml")
	flag.Parse()

	err := config.ParseFile(c)
	if err != nil {
		logrus.Panic(err)
	}
}

func main() {
	server := http.NewServer(http.Config{
		IP:   config.Http.IP,
		Port: config.Http.Port,
	})

	authenticator.SetUsers(config.WebDav.Users)
	dav := webdav.NewWebDav(webdav.Config{
		Prefix:        config.WebDav.Prefix,
		RootDir:       config.WebDav.RootDir,
		Authenticator: authenticator,
	})

	if err := dav.Init(); err != nil {
		logrus.WithFields(logrus.Fields{
			"service": "webdav",
		}).Panic(err)
	}
	server.AddHandlerFunc(config.WebDav.Prefix, dav.HandlerFunc)
	server.Init()

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		server.Serve()
	}()

	waitSignal()
	server.Shutdown(nil)
	wg.Wait()
}

func waitSignal() {
	stopChan := make(chan os.Signal)
	signal.Notify(stopChan, os.Interrupt)
	<-stopChan // wait for SIGINT
}
