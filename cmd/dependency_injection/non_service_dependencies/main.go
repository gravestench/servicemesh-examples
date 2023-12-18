package main

import (
	"github.com/gravestench/servicemesh"
)

func main() {
	mesh := servicemesh.New()

	// each service has a dependency that is not
	// actually resolved through the service mesh but by
	// some other means (that part is up to you).
	mesh.Add(newServiceWithAsyncDependencyResolution())
	mesh.Add(newServiceWithAsyncDependencyResolution())
	mesh.Add(newServiceWithAsyncDependencyResolution())
	mesh.Add(newServiceWithAsyncDependencyResolution())
	mesh.Add(newServiceWithAsyncDependencyResolution())

	mesh.Run()
}
