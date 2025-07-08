package app

import (
	"chat_with_me/app/web/middleware"
	"chat_with_me/app/web/service/api"
	"chat_with_me/app/web/service/ws"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type App struct {
	IndexController *api.IndexController
	wsController    *ws.WsController
	Engine          *gin.Engine
}

func NewApp(engine *gin.Engine, logger *zap.Logger, indexController *api.IndexController, wsController *ws.WsController) *App {
	//设置信任的代理
	if err := engine.SetTrustedProxies([]string{}); err != nil {
		panic(err)
	}
	//中间件
	engine.Use(middleware.LogsMiddleware(), middleware.AuthMiddleware(logger))
	wsGroup := engine.Group("/ws")
	{
		wsGroup.GET("/connect", wsController.Connect)
	}
	apiGroup := engine.Group("/api")
	{
		apiGroup.GET("/home", indexController.Home)
		apiGroup.GET("/create", indexController.Create)
		apiGroup.GET("/update/:id", indexController.Update)
		apiGroup.GET("/update_attr/:id", indexController.UpdateAttr)
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
