package main

import (
	"time"

	"github.com/gravestench/servicemesh"
)

func main() {
	mesh := servicemesh.New()

	for _, service := range []servicemesh.Service{
		&example{name: "foo"},
		&example{name: "bar"},
		&example{name: "baz"},
	} {
		mesh.Add(service)
	}

	go func() {
		time.Sleep(time.Second * 3)
		mesh.Shutdown()
	}()

	mesh.Run()
}
