package main

import (
	"github.com/faiface/mainthread"
	"github.com/gravestench/servicemesh"

	"github.com/gravestench/servicemesh-examples/services/config_file"
	"github.com/gravestench/servicemesh-examples/services/raylib_renderer"
)

func main() {
	mesh := servicemesh.New()
	r := &raylib_renderer.Service{}

	mesh.Add(&config_file.Service{RootDirectory: "~/.config/servicemesh/examples/simple_gui"})
	mesh.Add(r)

	// create 100 layers, each will show a moving circle
	for i := 0; i < 100; i++ {
		r.AddLayer(newLayer())
	}

	mainthread.Run(mesh.Run)
}
