package vms

import "github.com/google/wire"

var Providers = wire.NewSet(CreateContainer, MainLoader)
