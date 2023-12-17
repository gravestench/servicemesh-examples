package twitch_auth_code_grant_flow

import (
	"fmt"
	"log/slog"
	"sync"

	"github.com/gravestench/servicemesh"
	"github.com/gravestench/servicemesh-examples/services/config_file"
)

type Service struct {
	logger *slog.Logger
	cfg    config_file.Dependency
	mux    sync.Mutex

	stateString string

	token string
}

func (s *Service) Init(mesh servicemesh.Mesh) {

}

func (s *Service) Name() string {
	return "Twitch Integration (Auth code grant flow)"
}

func (s *Service) SetLogger(logger *slog.Logger) {
	s.logger = logger
}

func (s *Service) Logger() *slog.Logger {
	return s.logger
}

func (s *Service) Token() (string, error) {
	if s.token == "" {
		return "", fmt.Errorf("not yet authorized")
	}

	return s.token, nil
}
