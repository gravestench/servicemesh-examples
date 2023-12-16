package main

import (
	"github.com/gravestench/servicemesh"

	"github.com/gravestench/servicemesh-examples/services/config_file"
	"github.com/gravestench/servicemesh-examples/services/text_to_speech"
)

func main() {
	mesh := servicemesh.New()

	mesh.Add(&config_file.Service{RootDirectory: "~/.config/servicemesh/examples/text_to_speech"})
	mesh.Add(&text_to_speech.Service{})
	mesh.Add(&exampleServiceThatUsesTts{})

	mesh.Run()
}
