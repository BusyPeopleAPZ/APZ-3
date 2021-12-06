//+build wireinject

package main

import (
	"github.com/BusyPeopleAPZ/APZ-3/server/vms"
	"github.com/google/wire"
)

func ComposeApiServer(port HttpPortNumber) (*VMSApiServer, error) {
	wire.Build(
		DatabaseConnection,
		vms.Providers,
		wire.Struct(new(VMSApiServer), "Port", "VMSHandler"), 
	)
	return nil, nil
}
