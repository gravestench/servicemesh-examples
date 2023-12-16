package main

import (
	"github.com/gravestench/servicesmesh-examples/cmd/dependency_injection/service_dependencies/serviceA"
	"github.com/gravestench/servicesmesh-examples/cmd/dependency_injection/service_dependencies/serviceB"
)

func main() {
	rt := servicemesh.New()

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
	// will use the runtime to find and assign a service that matches the
	// interface they are looking for.
	rt.Add(serviceA.New("serviceA instance"))
	rt.Add(serviceB.New("serviceB instance"))

	// when this runs, you will see the runtime initiates the dependency
	// resolution and the services will end up with their dependencies met.
	rt.Run()
}
