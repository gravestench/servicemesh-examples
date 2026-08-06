package main

import (
	"errors"
	"fmt"
	"time"

	"github.com/gravestench/servicemesh"
)

func main() {
	mesh := servicemesh.New("Dependency Failure Example")
	mesh.SetDependencyResolutionTimeout(100 * time.Millisecond)

	// Register the observer before the unresolved service so it receives the
	// failure event.
	mesh.Add(&failureObserver{}).Wait()
	mesh.Add(&unresolvedService{}).Wait()
	mesh.Shutdown().Wait()
}

type unresolvedService struct{}

func (*unresolvedService) Init(servicemesh.Mesh) {}
func (*unresolvedService) Name() string          { return "Unresolved Service" }
func (*unresolvedService) DependenciesResolved() bool {
	return false
}
func (*unresolvedService) ResolveDependencies([]servicemesh.Service) {}

type failureObserver struct{}

func (*failureObserver) Init(servicemesh.Mesh) {}
func (*failureObserver) Name() string          { return "Dependency Failure Observer" }
func (*failureObserver) OnDependencyResolutionFailed(service servicemesh.Service, err error) {
	if errors.Is(err, servicemesh.ErrDependencyResolutionTimeout) {
		fmt.Printf("%s failed: %v\n", service.Name(), err)
	}
}

var _ servicemesh.EventHandlerDependencyResolutionFailed = (*failureObserver)(nil)
