package main

import (
	"github.com/gravestench/servicesmesh-examples/services/config_file"
	"github.com/gravestench/servicesmesh-examples/services/text_to_speech"
)

func main() {
	rt := servicemesh.New()

	rt.Add(&config_file.Service{RootDirectory: "~/.config/runtime/example/text_to_speech"})
	rt.Add(&text_to_speech.Service{})
	rt.Add(&exampleServiceThatUsesTts{})

	rt.Run()
}
