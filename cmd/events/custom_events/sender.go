package main

import (
	"log/slog"
	"time"

	"github.com/gravestench/servicemesh"
)

type sender struct {
	logger *slog.Logger
}

func (s *sender) SetLogger(logger *slog.Logger) {
	s.logger = logger
}

func (s *sender) Logger() *slog.Logger {
	return s.logger
}

func (s *sender) Ready() bool { return true }

func (s *sender) Init(mesh servicemesh.Mesh) {
	s.logger.Info("emitting 'test' event in 3 seconds...")

	time.Sleep(time.Second * 3)

	mesh.Events().Emit("test", "foo", 1, 2.3, []int{4, 5}).Wait()
}

func (s *sender) Name() string {
	return "sender"
}
