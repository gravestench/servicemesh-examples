package twitch_soundboard

import (
	"github.com/gravestench/servicesmesh-examples/services/config_file"
	"github.com/gravestench/servicesmesh-examples/services/twitch_integration"
)

// Ensure that Service implements the required interfaces.
var (
	_ servicemesh.Service          = &Service{}
	_ servicemesh.HasLogger        = &Service{}
	_ servicemesh.HasDependencies  = &Service{}
	_ config_file.HasDefaultConfig = &Service{}
	_ IsTwitchSoundboard           = &Service{}
)

type IsTwitchSoundboard interface {
	twitch_integration.OnPrivateMessage
}
