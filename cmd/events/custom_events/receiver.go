package main

import (
	"log/slog"

	"github.com/gravestench/servicemesh"
)

type receiver struct {
	logger *slog.Logger
}

func (r *receiver) SetLogger(logger *slog.Logger) {
	r.logger = logger
}

func (r *receiver) Logger() *slog.Logger {
	return r.logger
}

func (r *receiver) Init(mesh servicemesh.M) {
	mesh.Events().On("test", func(args ...any) {
		r.logger.Info("event received", "event", "test", "args", args)
	})
}

func (r *receiver) Name() string {
	return "receiver"
}
