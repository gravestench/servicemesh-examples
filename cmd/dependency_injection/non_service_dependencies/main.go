package main

import (
	"github.com/gravestench/servicemesh"

	"github.com/gravestench/servicemesh-examples/internal/examplemesh"
)

func main() {
	mesh := servicemesh.New()

	// each service has a dependency that is not
	// actually resolved through the service mesh but by
	// some other means (that part is up to you).
	examplemesh.AddAndWait(mesh,
		newServiceWithAsyncDependencyResolution(),
		newServiceWithAsyncDependencyResolution(),
		newServiceWithAsyncDependencyResolution(),
		newServiceWithAsyncDependencyResolution(),
		newServiceWithAsyncDependencyResolution(),
	)

	mesh.Run()
}
