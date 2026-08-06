package main

import (
	"github.com/gravestench/servicemesh"

	"github.com/gravestench/servicemesh-examples/internal/examplemesh"
)

func main() {
	mesh := servicemesh.New()

	examplemesh.AddAndWait(mesh, &sender{}, &receiver{})

	mesh.Run()
}
