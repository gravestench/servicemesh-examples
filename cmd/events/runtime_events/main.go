package main

import (
	"time"

	"github.com/gravestench/servicemesh"
)

func main() {
	mesh := servicemesh.New()

	go func() {
		time.Sleep(time.Second)
		mesh.Shutdown()
	}()

	mesh.Add(&listensForNewServices{})
	mesh.Add(&exampleService{name: "foo"})
	mesh.Add(&exampleService{name: "bar"})
	mesh.Add(&exampleService{name: "baz"})

	mesh.Run()
}
