package main

import (
	"github.com/gravestench/servicemesh"

	"github.com/gravestench/servicemesh-examples/services/config_file"
	"github.com/gravestench/servicemesh-examples/services/twitch_auth_code_grant_flow"
	"github.com/gravestench/servicemesh-examples/services/web_router"
	"github.com/gravestench/servicemesh-examples/services/web_server"
)

func main() {
	mesh := servicemesh.New()

	mesh.Add(&config_file.Service{RootDirectory: "~/.config"})
	mesh.Add(&web_server.Service{})
	mesh.Add(&web_router.Service{})
	mesh.Add(&twitch_auth_code_grant_flow.Service{})

	mesh.Run()
}
