package desktop_notification

import (
	"log/slog"
	"time"

	"github.com/gen2brain/beeep"
	"github.com/gravestench/servicemesh"

	"github.com/gravestench/servicemesh-examples/services/config_file"
)

type Service struct {
	logger     *slog.Logger
	cfgManager config_file.Manager
}

func (s *Service) Init(mesh servicemesh.Mesh) {

}

func (s *Service) Name() string {
	return "Desktop Notifications"
}

func (s *Service) SetLogger(logger *slog.Logger) {
	s.logger = logger
}

func (s *Service) Logger() *slog.Logger {
	return s.logger
}

func (s *Service) DependenciesResolved() bool {
	return s.cfgManager != nil
}

func (s *Service) ResolveDependencies(mesh servicemesh.Mesh) {
	for _, candidate := range mesh.Services() {
		if service, ok := candidate.(config_file.Manager); ok {
			s.cfgManager = service
		}
	}
}

func (s *Service) ConfigFilePath() string {
	return "desktop_notifications.json"
}

func (s *Service) Config() (*config_file.Config, error) {
	return s.cfgManager.GetConfig(s.ConfigFilePath())
}

func (s *Service) DefaultConfig() (cfg config_file.Config) {
	g := cfg.Group("Notifications")

	g.SetDefault("timeout", time.Second*5)

	return cfg
}

func (s *Service) Notify(title, message, appIcon string) {
	err := beeep.Notify(title, message, appIcon)
	if err != nil {
		s.logger.Error("sending notification", "error", err)
	}
}

func (s *Service) Alert(title, message, appIcon string) {
	err := beeep.Alert(title, message, appIcon)
	if err != nil {
		s.logger.Error("sending notification", "error", err)
	}
}
