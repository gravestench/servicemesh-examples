package main

import (
	"github.com/gravestench/servicemesh"

	"github.com/gravestench/servicemesh-examples/internal/examplemesh"
	"github.com/gravestench/servicemesh-examples/services/config_file"
	"github.com/gravestench/servicemesh-examples/services/twitch_auth_code_grant_flow"
	"github.com/gravestench/servicemesh-examples/services/web_router"
	"github.com/gravestench/servicemesh-examples/services/web_server"
)

func main() {
	mesh := servicemesh.New()

	examplemesh.AddAndWait(mesh,
		&config_file.Service{RootDirectory: "~/.config"},
		&web_server.Service{},
		&web_router.Service{},
		&twitch_auth_code_grant_flow.Service{},
	)

	mesh.Run()
}
