package raylib_renderer

import (
	"github.com/gravestench/servicemesh-examples/services/config_file"
)

const (
	keyGroupWindowSettings = "Window Settings"
	keyScreenWidth         = "Width"
	keyScreenHeight        = "Height"
)

func (s *Service) ConfigFilePath() string {
	return "raylib_renderer.json"
}

func (s *Service) Config() (*config_file.Config, error) {
	return s.cfgManager.GetConfig(s.ConfigFilePath())
}

func (s *Service) DefaultConfig() (cfg config_file.Config) {
	group := cfg.Group(keyGroupWindowSettings)

	group.SetDefault(keyScreenWidth, 1024)
	group.SetDefault(keyScreenHeight, 768)

	return cfg
}
