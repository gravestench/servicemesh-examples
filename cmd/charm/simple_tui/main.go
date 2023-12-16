package main

import (
	"os"
)

func main() {
	rt := servicemesh.New()

	logFile, err := os.OpenFile("/tmp/log.out", os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		panic(err)
	}

	rt.SetLogDestination(logFile)

	rt.Add(&bubbleteaService{})
	rt.Add(&filePickerService{})

	rt.Run()
}
