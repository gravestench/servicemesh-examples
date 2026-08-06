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

// Lifecycle events can be observed by implementing their corresponding
// optional handler interface.
func (s *listensForNewServices) OnServiceAdded(service servicemesh.Service) {
	if service == s {
		return
	}

	s.logger.Info("found a service", "found service", service.Name())
}

var _ servicemesh.EventHandlerServiceAdded = (*listensForNewServices)(nil)
