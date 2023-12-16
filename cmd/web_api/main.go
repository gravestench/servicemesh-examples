package main

import (
	"github.com/gravestench/servicemesh"

	"github.com/gravestench/servicemesh-examples/services/config_file"
	"github.com/gravestench/servicemesh-examples/services/web_router"
	"github.com/gravestench/servicemesh-examples/services/web_server"
)

func main() {
	mesh := servicemesh.New()

	// will manage the config files for the other services
	mesh.Add(&config_file.Service{})
	mesh.Add(&web_server.Service{})
	mesh.Add(&web_router.Service{})

	mesh.Add(&exampleRouteInitializer{}) // our example service that has routes

	mesh.Run()
}
