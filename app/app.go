package app

import (
	"chat_with_me/app/web/api"
	"chat_with_me/app/web/ws"
	"github.com/gin-gonic/gin"
)

func InitApp() *gin.Engine {
	wsController := ws.NewWsController()
	indexController := api.NewIndex()
	r := gin.Default()
	r.GET("/", indexController.Home)
	wsGroup := r.Group("/ws")
	//apiGroup := r.Group("/api")
	//建立ws链接
	wsGroup.GET("/connect", wsController.Connect)
	return r
}

func Run(r *gin.Engine) error {
	return r.Run(":8080")
}
