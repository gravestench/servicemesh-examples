package main

import (
	"log/slog"

	"github.com/gravestench/servicemesh"
)

func main() {
	mesh := servicemesh.New()

	mesh.Add(&service{})

	mesh.Run()
}

type service struct {
	logger *slog.Logger
}

func (s *service) Init(mesh servicemesh.Mesh) {
	mesh.SetLogLevel(int(slog.LevelWarn))
	s.logger.Debug("you should not see this")
	s.logger.Info("you should not see this")

	mesh.SetLogLevel(int(slog.LevelInfo))
	s.logger.Info("you should see this")
	s.logger.Warn("you should see this")
	s.logger.Error("you should see this")
}

func (s *service) Name() string {
	return "Example"
}

func (s *service) SetLogger(logger *slog.Logger) {
	s.logger = logger
}

func (s *service) Logger() *slog.Logger {
	return s.logger
}
