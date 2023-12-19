package main

import (
	"log/slog"
	"os"

	"github.com/gravestench/servicemesh"
)

func main() {
	mesh := servicemesh.New()
	mesh.SetLogHandler(slog.NewJSONHandler(os.Stdout, nil))

	mesh.Add(&example{name: "foo"})

	mesh.Run()
}
