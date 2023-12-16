package main

import (
	"time"
)

func main() {
	rt := servicemesh.New()

	for _, service := range []servicemesh.Service{
		&example{name: "foo"},
		&example{name: "bar"},
		&example{name: "baz"},
	} {
		rt.Add(service)
	}

	go func() {
		time.Sleep(time.Second * 3)
		rt.Shutdown()
	}()

	rt.Run()
}
