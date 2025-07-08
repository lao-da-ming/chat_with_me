//go:build wireinject
// +build wireinject

package cmd

import (
	"chat_with_me/app/web/app"
	"chat_with_me/app/web/log"
	"chat_with_me/app/web/service/api"
	"chat_with_me/app/web/service/ws"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

// wireApp init kratos application.
func WireApp(engine *gin.Engine) *app.App {
	panic(wire.Build(
		log.ProviderSet,
		api.ProviderSet,
		ws.ProviderSet,
		app.ProviderSet,
	))
}
