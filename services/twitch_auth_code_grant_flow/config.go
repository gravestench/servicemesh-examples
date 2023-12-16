package twitch_auth_code_grant_flow

import (
	"fmt"

	"github.com/gravestench/servicemesh-examples/services/config_file"
)

func (s *Service) ConfigFilePath() string {
	return "twitch_auth_code_grant_flow.json"
}

func (s *Service) Config() (*config_file.Config, error) {
	if s.cfg == nil {
		return nil, fmt.Errorf("no config manager")
	}

	return s.cfg.GetConfig(s.ConfigFilePath())
}

func (s *Service) DefaultConfig() (cfg config_file.Config) {
	g := cfg.Group(s.Name())

	for k, v := range map[string]any{
		"ClientID":    "",
		"ForceVerify": false,
		"RedirectURI": "http://localhost:8080/callback",
		"Scope":       "",
	} {
		g.Set(k, v)
	}

	return
}
