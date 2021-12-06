package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
)

type HttpPortNumber int

type VMSApiServer struct {
	Port HttpPortNumber
	VMSHandler http.HandlerFunc
	server *http.Server
}

func (s *VMSApiServer) Start() error {
	if s.VMSHandler == nil {
		return fmt.Errorf("VMS HTTP handler is not defined - cannot start")
	}
	if s.Port == 0 {
		return fmt.Errorf("port is not defined")
	}

	handler := http.NewServeMux()
	handler.HandleFunc("/vms", s.VMSHandler)

	s.server = &http.Server{
		Addr: fmt.Sprintf(":%d", s.Port),
		Handler: handler,
	}

	return s.server.ListenAndServe()
}

func (s *VMSApiServer) Stop() error {
	if s.server == nil {
		return errors.New("error: server has already been stopped")
	}

	return s.server.Shutdown(context.Background())
}
