package twitch_auth_code_grant_flow

import (
	"github.com/gravestench/servicemesh"

	"github.com/gravestench/servicemesh-examples/services/config_file"
	"github.com/gravestench/servicemesh-examples/services/web_router"
)

var _ servicemesh.Service = &Service{}
var _ servicemesh.HasLogger = &Service{}
var _ servicemesh.HasDependencies = &Service{}
var _ config_file.HasDefaultConfig = &Service{}
var _ web_router.IsRouteInitializer = &Service{}
var _ TwitchAuthGrantFlow = &Service{}

type TwitchAuthGrantFlow interface {
	Token() (string, error)
}
