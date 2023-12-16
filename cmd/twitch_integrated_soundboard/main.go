package main

import (
	"github.com/gravestench/servicemesh"

	"github.com/gravestench/servicemesh-examples/services/config_file"
	"github.com/gravestench/servicemesh-examples/services/twitch_integration"
	"github.com/gravestench/servicemesh-examples/services/twitch_soundboard"
)

func main() {
	mesh := servicemesh.New()

	// will manage the config files for the other services
	mesh.Add(&config_file.Service{})

	// This service has a dependency on the config manager
	mesh.Add(&twitch_integration.Service{})

	// this service has methods that satisfy an interface the twitch integration
	// service is looking for to bind event handlers.
	// see examples/services/twitch_integration/abstract.go for the interfaces
	mesh.Add(&twitch_soundboard.Service{})

	mesh.Run()
}
