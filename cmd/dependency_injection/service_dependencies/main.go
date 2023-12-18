package main

import (
	"github.com/gravestench/servicemesh"

	"github.com/gravestench/servicemesh-examples/cmd/dependency_injection/service_dependencies/serviceA"
	"github.com/gravestench/servicemesh-examples/cmd/dependency_injection/service_dependencies/serviceB"
)

func main() {
	mesh := servicemesh.New()

	// In this example we have two services which both rely on each other.
	//
	// Normally, this kind of dependency would result in cyclical import errors
	// and would require an interface be defined.
	//
	// serviceA and serviceB have methods A() and B(), respectively.
	//
	// serviceA has a dependency on serviceB and will call the B() method.
	// the inverse is true of serviceB, which needs to call the A() method.
	//
	// both of these services implement servicemesh.HasDependencies, and they both
	// will use the service mesh to find and assign a service that matches the
	// interface they are looking for.
	mesh.Add(serviceA.New("serviceA instance"))
	mesh.Add(serviceB.New("serviceB instance"))

	// when this runs, you will see the service mesh initiates the dependency
	// resolution and the services will end up with their dependencies met.
	mesh.Run()
}
