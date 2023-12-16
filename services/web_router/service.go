package web_router

import (
	"time"

	"github.com/gin-gonic/gin"

	"github.com/gravestench/servicesmesh-examples/services/config_file"
	"github.com/gravestench/servicesmesh-examples/services/web_router/middleware/static_assets"
)

type Service struct {
	log        *slog.Logger
	cfgManager config_file.Dependency

	root *gin.Engine

	boundServices map[string]*struct{} // holds 0-size entries

	config struct {
		debug bool
	}

	reloadDebounce time.Time
}

func (s *Service) SetLogger(l *slog.Logger) {
	s.log = l
}

func (s *Service) Logger() *slog.Logger {
	return s.log
}

func (s *Service) Init(mesh servicemesh.M) {
	gin.SetMode("release")
	rt.Add(&static_assets.Middleware{})
	s.root = gin.New()
	go s.beginDynamicRouteBinding(rt)
}

func (s *Service) Name() string {
	return "Web Router"
}
