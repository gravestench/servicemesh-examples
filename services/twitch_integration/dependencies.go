package twitch_integration

import (
	"github.com/gravestench/servicesmesh-examples/services/config_file"
)

func (s *Service) DependenciesResolved() bool {
	if s.cfgManager == nil {
		return false
	}

	if cfg, err := s.Config(); cfg == nil || err != nil {
		return false
	}

	return true
}

func (s *Service) ResolveDependencies(mesh servicemesh.M) {
	for _, service := range rt.Services() {
		if candidate, ok := service.(config_file.Dependency); ok {
			s.cfgManager = candidate
		}
	}
}
