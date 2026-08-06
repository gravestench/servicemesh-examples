package main

import (
	"github.com/gravestench/servicemesh"

	"github.com/gravestench/servicemesh-examples/internal/examplemesh"
	"github.com/gravestench/servicemesh-examples/services/config_file"
	"github.com/gravestench/servicemesh-examples/services/text_to_speech"
)

func main() {
	mesh := servicemesh.New()

	examplemesh.AddAndWait(mesh,
		&config_file.Service{RootDirectory: "~/.config/servicemesh/examples/text_to_speech"},
		&text_to_speech.Service{},
		&exampleServiceThatUsesTts{},
	)

	mesh.Run()
}
