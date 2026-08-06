package main

import (
	"time"

	"github.com/gravestench/servicemesh"

	"github.com/gravestench/servicemesh-examples/internal/examplemesh"
)

func main() {
	mesh := servicemesh.New()

	go func() {
		time.Sleep(time.Second)
		mesh.Shutdown().Wait()
	}()

	// Register the listener first so it observes the remaining additions.
	mesh.Add(&listensForNewServices{}).Wait()
	examplemesh.AddAndWait(mesh,
		&exampleService{name: "foo"},
		&exampleService{name: "bar"},
		&exampleService{name: "baz"},
	)

	mesh.Run()
}
