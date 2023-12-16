package raylib_renderer

import (
	"github.com/gravestench/servicemesh"

	"github.com/gravestench/servicemesh-examples/services/config_file"
)

// Ensure that Service implements the required interfaces.
var (
	_ servicemesh.Service             = &Service{}
	_ servicemesh.HasLogger           = &Service{}
	_ servicemesh.HasDependencies     = &Service{}
	_ servicemesh.HasGracefulShutdown = &Service{}
	_ config_file.HasDefaultConfig    = &Service{}
	_ LayerRenderer                   = &Service{}
)

type LayerRenderer interface {
	AddLayer(RenderableLayer)
}

type RenderableLayer interface {
	OnRender()
}
