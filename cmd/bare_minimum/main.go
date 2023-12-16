package main

import (
	"github.com/gravestench/servicemesh"
)

func main() {
	mesh := servicemesh.New()

	mesh.Add(&example{name: "foo"})

	mesh.Run()
}
