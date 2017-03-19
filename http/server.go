package http

import (
	"fmt"
	"net/http"

	"github.com/Sirupsen/logrus"
)

type Server struct {
	config Config
	mux    *http.ServeMux
}

func NewServer(config Config) *Server {
	return &Server{
		config: config,
		mux:    http.DefaultServeMux,
	}
}

func (s *Server) AddHandlerFunc(path string, handlerFunc http.HandlerFunc) {
	s.mux.HandleFunc(path, handlerFunc)
}

func (s *Server) Serve() (err error) {
	logrus.WithFields(logrus.Fields{
		"service": "web",
		"ip":      s.config.IP,
		"port":    s.config.Port,
	}).Info("Serve")

	listenAddr := fmt.Sprintf("%s:%d", s.config.IP, s.config.Port)

	err = http.ListenAndServe(listenAddr, s.mux)

	if err != nil {
		logrus.WithFields(logrus.Fields{
			"service": "web",
			"ip":      s.config.IP,
			"port":    s.config.Port,
		}).Error("Can't listen")
		return
	}

	logrus.WithFields(logrus.Fields{
		"service": "web",
	}).Info("Shutdown")
	return
}
