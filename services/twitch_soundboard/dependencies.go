package twitch_soundboard

import (
	"github.com/gravestench/servicemesh"

	"github.com/gravestench/servicemesh-examples/services/config_file"
	"github.com/gravestench/servicemesh-examples/services/desktop_notification"
)

func (s *Service) ResolveDependencies(services []servicemesh.Service) {
	for _, service := range services {
		if instance, ok := service.(config_file.Manager); ok {
			s.configManager = instance
		}

		if instance, ok := service.(desktop_notification.SendsNotifications); ok {
			s.notification = instance
		}
	}
}

func (s *Service) DependenciesResolved() bool {
	if s.configManager == nil {
		return false
	}

	return true
}
