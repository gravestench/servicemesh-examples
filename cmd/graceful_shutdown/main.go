package main

import (
	"time"

	"github.com/gravestench/servicemesh"

	"github.com/gravestench/servicemesh-examples/internal/examplemesh"
)

func main() {
	mesh := servicemesh.New()

	services := []servicemesh.Service{
		&example{name: "foo"},
		&example{name: "bar"},
		&example{name: "baz"},
	}
	examplemesh.AddAndWait(mesh, services...)

	go func() {
		time.Sleep(time.Second * 3)
		mesh.Shutdown().Wait()
	}()

	mesh.Run()
}
