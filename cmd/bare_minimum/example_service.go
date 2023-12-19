package main

import (
	"github.com/gravestench/servicemesh"
)

type example struct {
	name string
}

func (e *example) Init(mesh servicemesh.Mesh) {
	return
}

func (e *example) Name() string {
	return e.name
}

func (s *example) Ready() bool { return true }
