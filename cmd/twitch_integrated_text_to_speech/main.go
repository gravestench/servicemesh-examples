package main

import (
	"time"

	"github.com/gravestench/servicemesh"

	"github.com/gravestench/servicemesh-examples/internal/examplemesh"
	"github.com/gravestench/servicemesh-examples/services/config_file"
	"github.com/gravestench/servicemesh-examples/services/text_to_speech"
	"github.com/gravestench/servicemesh-examples/services/twitch_integration"
)

func main() {
	mesh := servicemesh.New()

	cfgDir := "~/.config/servicemesh/examples/twitch_integrated_text_to_speech"

	examplemesh.AddAndWait(mesh,
		&config_file.Service{RootDirectory: cfgDir},
		&twitch_integration.Service{},
		&text_to_speech.Service{},
		// Connects the Twitch integration to text-to-speech.
		&glueService{startupTime: time.Now(), onJoinDelay: time.Minute},
	)

	mesh.Run()
}
