package main

import (
	"github.com/gravestench/servicemesh"

	"github.com/gravestench/servicemesh-examples/internal/examplemesh"
	"github.com/gravestench/servicemesh-examples/services/config_file"
)

func main() {
	mesh := servicemesh.New()
	cfgManager := &config_file.Service{}

	examplemesh.AddAndWait(mesh,
		cfgManager,
		// This service has a dependency on the config manager.
		&serviceThatUsesConfigManager{},
	)

	mesh.Run()
}
