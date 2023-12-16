package main

import (
	"os"

	"github.com/gravestench/servicemesh"
)

func main() {
	mesh := servicemesh.New()

	logFile, err := os.OpenFile("/tmp/log.out", os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		panic(err)
	}

	mesh.SetLogDestination(logFile)

	mesh.Add(&bubbleteaService{})
	mesh.Add(&filePickerService{})

	mesh.Run()
}
