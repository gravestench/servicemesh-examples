package main

import (
	"github.com/gravestench/servicemesh"

	"github.com/gravestench/servicemesh-examples/services/config_file"
)

func main() {
	mesh := servicemesh.New()
	cfgManager := &config_file.Service{}

	mesh.Add(cfgManager)

	// This service has a dependency on the config manager
	mesh.Add(&serviceThatUsesConfigManager{})

	mesh.Run()
}
