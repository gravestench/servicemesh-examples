package main

import (
	"time"
)

func main() {
	rt := servicemesh.New()

	go func() {
		time.Sleep(time.Second)
		rt.Shutdown()
	}()

	rt.Add(&listensForNewServices{})
	rt.Add(&exampleService{name: "foo"})
	rt.Add(&exampleService{name: "bar"})
	rt.Add(&exampleService{name: "baz"})

	rt.Run()
}
