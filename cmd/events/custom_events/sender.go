package main

import (
	"time"
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

func (s *sender) Init(mesh servicemesh.M) {
	s.logger.Info().Msgf("emitting event in 3 seconds...")

	time.Sleep(time.Second * 3)

	rt.Events().Emit("test", "foo", 1, 2.3, []int{4, 5}).Wait()
}

func (s *sender) Name() string {
	return "sender"
}
