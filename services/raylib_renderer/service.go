package raylib_renderer

import (
	"fmt"
	"log/slog"
	"sync"
	"unsafe"

	"github.com/faiface/mainthread"
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/gravestench/servicemesh"
	"github.com/gravestench/servicesmesh-examples/services/config_file"
)

type Service struct {
	log        *slog.Logger
	mux        sync.Mutex
	cfgManager config_file.Manager // dependency
	layers     map[string]RenderableLayer
	order      []string
}

func (s *Service) Init(mesh servicemesh.M) {
	defer func() { _ = recover() /* don't worry about it */ }()

	if s.layers == nil {
		s.layers = make(map[string]RenderableLayer)
	}

	// raylib requires servicemesh.Run() to be passed into mainthread.Run in main.go
	go s.gatherLayers(mesh)
	s.initRenderer()
	s.renderServicesAsLayers(mesh)
}

func (s *Service) OnShutdown() {
	mainthread.Run(rl.CloseWindow)
}

func (s *Service) AddLayer(layer RenderableLayer) {
	s.mux.Lock()
	defer s.mux.Unlock()

	if s.layers == nil {
		s.layers = make(map[string]RenderableLayer)
	}

	layerName := pointerAsString(layer)

	if s.layers[layerName] != nil {
		return
	}

	s.layers[layerName] = layer
	s.order = append(s.order, layerName)
}

func (s *Service) Name() string {
	return "Raylib Renderer"
}

func (s *Service) SetLogger(logger *slog.Logger) {
	s.log = logger
}

func (s *Service) Logger() *slog.Logger {
	return s.log
}

func pointerAsString(ptr interface{}) string {
	return fmt.Sprintf("%p", unsafe.Pointer(&ptr))
}
