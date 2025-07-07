package app

import (
	"chat_with_me/app/web/service/api"
	"chat_with_me/app/web/service/ws"
	"github.com/gin-gonic/gin"
)

type App struct {
	IndexController *api.IndexController
	wsController    *ws.WsController
	Engine          *gin.Engine
}

func NewApp(engine *gin.Engine, indexController *api.IndexController, wsController *ws.WsController) *App {
	//设置信任的代理
	if err := engine.SetTrustedProxies([]string{}); err != nil {
		panic(err)
	}
	wsGroup := engine.Group("/ws")
	{
		wsGroup.GET("/connect", wsController.Connect)
	}
	apiGroup := engine.Group("/api")
	{
		apiGroup.GET("/home", indexController.Home)
	}
	return &App{
		IndexController: indexController,
		wsController:    wsController,
		Engine:          engine,
	}
}

func (app *App) Run() error {
	return app.Engine.Run(":8080")
}
