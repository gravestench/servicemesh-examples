package main

import (
	"github.com/rs/zerolog"
)

func main() {
	rt := servicemesh.New()

	rt.Add(&service{})

	rt.Run()
}

type service struct {
	logger *slog.Logger
}

func (s *service) Init(mesh servicemesh.M) {
	rt.SetLogLevel(zerolog.WarnLevel)
	s.logger.Trace().Msg("you should not see this")
	s.logger.Debug().Msg("you should not see this")
	s.logger.Info().Msg("you should not see this")

	rt.SetLogLevel(zerolog.InfoLevel)
	s.logger.Info().Msg("you should see this")
	s.logger.Debug().Msg("you should not see this")
	s.logger.Warn().Msg("you should see this")
	s.logger.Error().Msg("you should see this")

	rt.SetLogLevel(zerolog.TraceLevel)
	s.logger.Trace().Msg("you should see this")
	s.logger.Fatal().Msg("you should see this")
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
