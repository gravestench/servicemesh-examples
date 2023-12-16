package main

import (
	"github.com/gravestench/servicesmesh-examples/services/config_file"
)

func main() {
	rt := servicemesh.New()
	cfgManager := &config_file.Service{}

	rt.Add(cfgManager)

	// This service has a dependency on the config manager
	rt.Add(&serviceThatUsesConfigManager{})

	rt.Run()
}
