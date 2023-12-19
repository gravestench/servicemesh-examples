package web_server

import (
	"github.com/gravestench/servicemesh"

	"github.com/gravestench/servicemesh-examples/services/config_file"
	"github.com/gravestench/servicemesh-examples/services/web_router"
)

func (s *Service) DependenciesResolved() bool {
	if s.router == nil {
		return false
	}

	if s.cfgManager == nil {
		return false
	}

	return true
}

func (s *Service) ResolveDependencies(services []servicemesh.Service) {
	for _, other := range services {
		if router, ok := other.(web_router.IsWebRouter); ok {
			if router.RouteRoot() != nil {
				s.router = router
			}
		}

		if cfg, ok := other.(config_file.Manager); ok && s.cfgManager == nil {
			s.cfgManager = cfg
		}
	}
}
