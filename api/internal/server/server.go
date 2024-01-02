package server

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

type Server interface {
	Handle(pattern string, handler http.Handler)
	ListenAndServe() error
	Shutdown() error
}

func New(port uint) Server {
	mux := &http.ServeMux{}
	srv := &http.Server{
		Addr:              fmt.Sprintf(":%d", port),
		Handler:           mux,
		MaxHeaderBytes:    2 * 1024 * 1024,
		IdleTimeout:       time.Second * 60,
		ReadHeaderTimeout: time.Second * 60,
		ReadTimeout:       time.Second * 60,
		WriteTimeout:      time.Second * 60,
	}
	srv.SetKeepAlivesEnabled(true)
	return &server{
		mux: mux,
		srv: srv,
	}
}

type server struct {
	mux *http.ServeMux
	srv *http.Server
}

func (s *server) Handle(pattern string, handler http.Handler) {
	s.mux.Handle(pattern, handler)
}

func (s *server) ListenAndServe() error {
	return s.srv.ListenAndServe()
}

func (s *server) Shutdown() error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	return s.srv.Shutdown(ctx)
}
