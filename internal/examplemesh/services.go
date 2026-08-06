// Package examplemesh contains small helpers shared by the example commands.
package examplemesh

import (
	"sync"

	"github.com/gravestench/servicemesh"
)

// AddAndWait registers every service before waiting for initialization. Adding
// the full set first is important when services depend on each other.
func AddAndWait(mesh servicemesh.Mesh, services ...servicemesh.Service) {
	waits := make([]*sync.WaitGroup, 0, len(services))
	for _, service := range services {
		waits = append(waits, mesh.Add(service))
	}
	for _, wait := range waits {
		wait.Wait()
	}
}
