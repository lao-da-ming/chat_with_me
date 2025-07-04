package ws

import (
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(NewWsController)
