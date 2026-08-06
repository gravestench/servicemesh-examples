package main

import (
	"os"

	"github.com/gravestench/servicemesh"

	"github.com/gravestench/servicemesh-examples/internal/examplemesh"
)

func main() {
	mesh := servicemesh.New()

	logFile, err := os.OpenFile("/tmp/log.out", os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		panic(err)
	}

	mesh.SetLogDestination(logFile)

	examplemesh.AddAndWait(mesh, &bubbleteaService{}, &filePickerService{})

	mesh.Run()
}
