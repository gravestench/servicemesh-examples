package main

import (
	"github.com/gravestench/servicemesh"

	"github.com/gravestench/servicemesh-examples/internal/examplemesh"
	"github.com/gravestench/servicemesh-examples/services/config_file"
	"github.com/gravestench/servicemesh-examples/services/twitch_integration"
	"github.com/gravestench/servicemesh-examples/services/twitch_soundboard"
)

func main() {
	mesh := servicemesh.New()

	examplemesh.AddAndWait(mesh,
		// Manages configuration files for the other services.
		&config_file.Service{},
		// Depends on the configuration manager.
		&twitch_integration.Service{},
		// Exposes handlers consumed by the Twitch integration.
		&twitch_soundboard.Service{},
	)

	mesh.Run()
}
