package http

import (
	"context"
	"fmt"
	"net/http"

	log "github.com/Sirupsen/logrus"
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
	log.WithFields(log.Fields{
		"ip":   s.config.IP,
		"port": s.config.Port,
	}).Info("init http server")

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
		log.WithFields(log.Fields{
			"ip":   s.config.IP,
			"port": s.config.Port,
		}).Info("http server stopped")
		return
	}

	log.Info("shutdown")
	return
}

func (s *Server) Shutdown(ctx context.Context) {
	s.server.Shutdown(ctx)
}
