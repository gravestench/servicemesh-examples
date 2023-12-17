package web_router

import (
	"github.com/gravestench/servicemesh"

	"github.com/gravestench/servicemesh-examples/services/config_file"
)

func (s *Service) DependenciesResolved() bool {
	if s.cfgManager == nil {
		return false
	}

	return true
}

func (s *Service) ResolveDependencies(mesh servicemesh.Mesh) {
	for _, other := range mesh.Services() {
		if cfg, ok := other.(config_file.Manager); ok {
			s.cfgManager = cfg
		}
	}
}
