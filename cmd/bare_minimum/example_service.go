package main

import (
	"github.com/gravestench/servicemesh"
)

type example struct {
	name string
}

func (e *example) Init(mesh servicemesh.M) {
	return
}

func (e *example) Name() string {
	return e.name
}
