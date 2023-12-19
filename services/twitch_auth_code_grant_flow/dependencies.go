package twitch_auth_code_grant_flow

import (
	"github.com/gravestench/servicemesh"

	"github.com/gravestench/servicemesh-examples/services/config_file"
)

func (s *Service) DependenciesResolved() bool {
	return s.cfg != nil
}

func (s *Service) ResolveDependencies(services []servicemesh.Service) {
	for _, service := range services {
		if candidate, ok := service.(config_file.Dependency); ok {
			s.cfg = candidate
		}
	}
}
