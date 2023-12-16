package desktop_notification

import (
	"github.com/gravestench/servicemesh"
	"github.com/gravestench/servicesmesh-examples/services/config_file"
)

// Ensure that Service implements the required interfaces.
var (
	_ servicemesh.Service          = &Service{}
	_ servicemesh.HasLogger        = &Service{}
	_ servicemesh.HasDependencies  = &Service{}
	_ config_file.HasDefaultConfig = &Service{}
	_ SendsNotifications           = &Service{}
)

type SendsNotifications interface {
	Notify(title, message, appIcon string)
	Alert(title, message, appIcon string)
}
