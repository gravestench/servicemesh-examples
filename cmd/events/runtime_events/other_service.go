package main

import (
	"github.com/gravestench/servicemesh"
)

// the listener service will react to instances of this being added

type exampleService struct {
	name string
}

func (e *exampleService) Init(mesh servicemesh.Mesh) {
	// noop
}

func (e *exampleService) Name() string {
	return e.name
}
