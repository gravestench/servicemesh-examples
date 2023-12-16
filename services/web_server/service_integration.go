package web_server

import (
	"github.com/gravestench/servicesmesh-examples/services/config_file"
)

var (
	_ servicemesh.Service          = &Service{}
	_ servicemesh.HasLogger        = &Service{}
	_ config_file.HasDefaultConfig = &Service{}
	_ IsWebServer                  = &Service{}
)

type Dependency = IsWebServer

type IsWebServer interface {
	RestartServer()
	StartServer()
	StopServer()
}
