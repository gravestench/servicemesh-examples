package main

import (
	"log/slog"

	"github.com/gravestench/servicemesh"
)

type listensForNewServices struct {
	logger *slog.Logger
}

func (s *listensForNewServices) SetLogger(logger *slog.Logger) {
	s.logger = logger
}

func (s *listensForNewServices) Logger() *slog.Logger {
	return s.logger
}

func (s *listensForNewServices) Init(mesh servicemesh.Mesh) {
	// noop
}

func (s *listensForNewServices) Name() string {
	return "listener"
}

// there are a bunch of runtime events to bind to via
// implementing an interface like this one

func (s *listensForNewServices) OnServiceAdded(args ...interface{}) {
	if len(args) < 1 {
		return
	}

	service, ok := args[0].(servicemesh.Service)
	if !ok {
		return
	}

	if service == s {
		return
	}

	s.logger.Info("found a service", "found service", service.Name())
}
