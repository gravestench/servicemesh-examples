package main

import (
	"time"

	"github.com/gravestench/servicemesh"

	"github.com/gravestench/servicemesh-examples/services/config_file"
	"github.com/gravestench/servicemesh-examples/services/text_to_speech"
	"github.com/gravestench/servicemesh-examples/services/twitch_integration"
)

func main() {
	mesh := servicemesh.New()

	cfgDir := "~/.config/servicemesh/examples/twitch_integrated_text_to_speech"

	mesh.Add(&config_file.Service{RootDirectory: cfgDir})
	mesh.Add(&twitch_integration.Service{})
	mesh.Add(&text_to_speech.Service{})
	mesh.Add(&glueService{startupTime: time.Now(), onJoinDelay: time.Second * 60}) // this connects the twitch integration to the TTS

	mesh.Run()
}
