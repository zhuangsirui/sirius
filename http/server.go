package http

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Sirupsen/logrus"
)

type Server struct {
	config Config
	mux    *http.ServeMux
	server *http.Server
}

func NewServer(config Config) *Server {
	return &Server{
		config: config,
		mux:    http.NewServeMux(),
	}
}

func (s *Server) Init() {
	logrus.WithFields(logrus.Fields{
		"service": "http",
		"ip":      s.config.IP,
		"port":    s.config.Port,
	}).Info("serve")

	s.server = &http.Server{
		Addr:    fmt.Sprintf("%s:%d", s.config.IP, s.config.Port),
		Handler: s.mux,
	}
}

func (s *Server) AddHandlerFunc(path string, handlerFunc http.HandlerFunc) {
	s.mux.HandleFunc(path, handlerFunc)
}

func (s *Server) Serve() {
	if err := s.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		logrus.WithFields(logrus.Fields{
			"service": "http",
			"ip":      s.config.IP,
			"port":    s.config.Port,
		}).Info("http server stopped")
		return
	}

	logrus.WithFields(logrus.Fields{
		"service": "web",
	}).Info("shutdown")
	return
}

func (s *Server) Shutdown(ctx context.Context) {
	s.server.Shutdown(ctx)
}
