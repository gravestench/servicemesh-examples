package main

import (
	"github.com/gravestench/servicemesh"

	"github.com/gravestench/servicemesh-examples/internal/examplemesh"
	"github.com/gravestench/servicemesh-examples/services/config_file"
	"github.com/gravestench/servicemesh-examples/services/web_router"
	"github.com/gravestench/servicemesh-examples/services/web_server"
)

func main() {
	mesh := servicemesh.New()

	examplemesh.AddAndWait(mesh,
		// Manages configuration files for the other services.
		&config_file.Service{},
		&web_server.Service{},
		&web_router.Service{},
		&exampleRouteInitializer{}, // Example service that contributes routes.
	)

	mesh.Run()
}
