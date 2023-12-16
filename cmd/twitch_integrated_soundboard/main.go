package main

import (
	"github.com/gravestench/servicesmesh-examples/services/config_file"
	"github.com/gravestench/servicesmesh-examples/services/twitch_integration"
	"github.com/gravestench/servicesmesh-examples/services/twitch_soundboard"
)

func main() {
	rt := servicemesh.New()

	// will manage the config files for the other services
	rt.Add(&config_file.Service{})

	// This service has a dependency on the config manager
	rt.Add(&twitch_integration.Service{})

	// this service has methods that satisfy an interface the twitch integration
	// service is looking for to bind event handlers.
	// see examples/services/twitch_integration/abstract.go for the interfaces
	rt.Add(&twitch_soundboard.Service{})

	rt.Run()
}
