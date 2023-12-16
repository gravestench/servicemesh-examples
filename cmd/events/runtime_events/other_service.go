package main

// the listener service will react to instances of this being added

type exampleService struct {
	name string
}

func (e *exampleService) Init(mesh servicemesh.M) {
	// noop
}

func (e *exampleService) Name() string {
	return e.name
}