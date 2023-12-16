package twitch_integration

import (
	"log/slog"

	"github.com/gempir/go-twitch-irc/v2"
	"github.com/gravestench/servicemesh"

	"github.com/gravestench/servicemesh-examples/services/config_file"
)

type Service struct {
	logger     *slog.Logger
	cfgManager config_file.Manager

	twitchIrcClient *twitch.Client
}

func (s *Service) Init(mesh servicemesh.M) {
	go s.setupClient()
	go s.loopBindHandlers(mesh)
}

func (s *Service) Name() string {
	return "Twitch Integration"
}

func (s *Service) SetLogger(logger *slog.Logger) {
	s.logger = logger
}

func (s *Service) Logger() *slog.Logger {
	return s.logger
}
