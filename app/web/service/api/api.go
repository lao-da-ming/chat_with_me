package api

import "github.com/google/wire"

var ProviderSet = wire.NewSet(NewIndexController)
