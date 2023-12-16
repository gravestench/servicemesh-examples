package web_server

import (
	"net/http"

	"github.com/gravestench/servicesmesh-examples/services/config_file"
	"github.com/gravestench/servicesmesh-examples/services/web_router"
)

type Service struct {
	log        *slog.Logger
	router     web_router.Dependency
	cfgManager config_file.Dependency
	server     *http.Server
	lastConfig string
}

func (s *Service) Init(mesh servicemesh.M) {
	s.StartServer()
}

func (s *Service) SetLogger(l *slog.Logger) {
	s.log = l
}

func (s *Service) Logger() *slog.Logger {
	return s.log
}

func (s *Service) Name() string {
	return "Web Server"
}
