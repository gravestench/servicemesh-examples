package text_to_speech

import (
	"fmt"
	"path/filepath"

	"github.com/gravestench/servicesmesh-examples/services/config_file"
)

func (s *Service) ConfigFilePath() string {
	return "text_to_speech.json"
}

func (s *Service) Config() (*config_file.Config, error) {
	if s.cfgManager == nil {
		return nil, fmt.Errorf("no config manager")
	}

	return s.cfgManager.GetConfig(s.ConfigFilePath())
}

func (s *Service) DefaultConfig() (cfg config_file.Config) {
	g := cfg.Group("Text to speech")

	cfgDir := s.cfgManager.ConfigDirectory()
	g.SetDefault("directory", filepath.Join(cfgDir, "audio_files"))
	g.SetDefault("mplayer-handler", false)

	return
}
