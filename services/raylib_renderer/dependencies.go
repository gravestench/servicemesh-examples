package raylib_renderer

import (
	"github.com/gravestench/servicemesh"
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
	for _, service := range mesh.Services() {
		if candidate, ok := service.(config_file.Manager); ok {
			s.cfgManager = candidate
		}
	}
}
