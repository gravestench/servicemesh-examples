package main

type example struct {
	name string
}

func (e *example) Init(r servicemesh.R) {
	return
}

func (e *example) Name() string {
	return e.name
}
