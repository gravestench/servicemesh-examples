package twitch_soundboard

import (
	"github.com/gravestench/servicesmesh-examples/services/config_file"
	"github.com/gravestench/servicesmesh-examples/services/desktop_notification"
)

// this is an example service that implements only the OnPrivateMessage handler

type Service struct {
	configManager config_file.Manager // dependency on config file manager
	notification  desktop_notification.SendsNotifications
	log           *slog.Logger
}

func (s *Service) Init(r servicemesh.R) {
	// nothing to do
}

func (s *Service) Name() string {
	return "Twitch Chat Soundboard"
}

func (s *Service) SetLogger(logger *slog.Logger) {
	s.log = logger
}

func (s *Service) Logger() *slog.Logger {
	return s.log
}
