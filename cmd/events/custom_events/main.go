package main

import (
	"github.com/gravestench/servicemesh"
)

func main() {
	mesh := servicemesh.New()

	mesh.Add(&sender{})
	mesh.Add(&receiver{})

	mesh.Run()
}
