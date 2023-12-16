package main

import (
	"github.com/gravestench/servicesmesh-examples/services/config_file"
)

var _ interface {
	servicemesh.S
	servicemesh.HasLogger
	servicemesh.HasDependencies
} = &serviceThatUsesConfigManager{}

type serviceThatUsesConfigManager struct {
	configManager config_file.Manager // dependency on config file manager
	log           *slog.Logger
}

func (s *serviceThatUsesConfigManager) ResolveDependencies(mesh servicemesh.M) {
	for _, service := range rt.Services() {
		if instance, ok := service.(config_file.Manager); ok {
			s.configManager = instance
		}
	}
}

func (s *serviceThatUsesConfigManager) DependenciesResolved() bool {
	return s.configManager != nil
}

func (s *serviceThatUsesConfigManager) Init(r servicemesh.R) {
	cfg, err := s.configManager.GetConfig("test.json")
	if err != nil {
		s.log.Fatal().Msgf("couldnt load example config file", "error", err)
	}

	group := cfg.Group("foo")

	group.Set("a", 1)
	group.Set("b", 2)
	group.Set("c", 3)

	s.configManager.SaveConfig("test.json")
}

func (s *serviceThatUsesConfigManager) Name() string {
	return "Config User"
}

func (s *serviceThatUsesConfigManager) SetLogger(logger *slog.Logger) {
	s.log = logger
}

func (s *serviceThatUsesConfigManager) Logger() *slog.Logger {
	return s.log
}
