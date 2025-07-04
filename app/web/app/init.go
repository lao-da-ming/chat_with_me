package app

import (
	"chat_with_me/app/web/api"
	"chat_with_me/app/web/ws"
	"github.com/gin-gonic/gin"
)

type App struct {
	IndexController *api.IndexController
	wsController    *ws.WsController
	Engine          *gin.Engine
}

func NewApp(engine *gin.Engine, indexController *api.IndexController, wsController *ws.WsController) *App {
	engine.GET("/ws/connect", wsController.Connect)
	engine.GET("/:key", indexController.Home)
	return &App{
		IndexController: indexController,
		wsController:    wsController,
		Engine:          engine,
	}
}

func (app *App) Run() error {
	return app.Engine.Run(":8080")
}
