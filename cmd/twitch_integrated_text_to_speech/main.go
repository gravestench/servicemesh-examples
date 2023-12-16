package main

import (
	"time"

	"github.com/gravestench/servicesmesh-examples/services/config_file"
	"github.com/gravestench/servicesmesh-examples/services/text_to_speech"
	"github.com/gravestench/servicesmesh-examples/services/twitch_integration"
)

func main() {
	rt := servicemesh.New()

	cfgDir := "~/.config/runtime/example/twitch_integrated_text_to_speech"

	rt.Add(&config_file.Service{RootDirectory: cfgDir})
	rt.Add(&twitch_integration.Service{})
	rt.Add(&text_to_speech.Service{})
	rt.Add(&glueService{startupTime: time.Now(), onJoinDelay: time.Second * 60}) // this connects the twitch integration to the TTS

	rt.Run()
}
