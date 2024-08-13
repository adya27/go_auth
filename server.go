package todo

import (
	"context"
	"net/http"
	"time"
)

type Server struct {
	htpServer *http.Server
}

func (s *Server) Run(port string, handler http.Handler) error {
	s.htpServer = &http.Server{
		Addr:           ":" + port,
		Handler:        handler,
		MaxHeaderBytes: 1 << 20,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second}

	return s.htpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.Shutdown(ctx)
}
