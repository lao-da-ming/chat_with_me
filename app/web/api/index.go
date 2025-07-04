package api

import (
	"chat_with_me/app/web/ws"
	"github.com/gin-gonic/gin"
)

type IndexController struct {
}

func NewIndexController() *IndexController {
	return &IndexController{}
}
func (i *IndexController) Home(c *gin.Context) {
	conn, ok := ws.ConnectPool.Load(c.Param("key"))
	if !ok {
		c.JSON(200, gin.H{"msg": "链接不存在"})
	}
	fd, ok := conn.(*ws.ConnObj)
	if !ok {
		c.JSON(200, gin.H{"msg": "断言链接失败"})
	}
	fd.WriteToSendChan([]byte("9999999999999999999999999"))
	c.JSON(200, gin.H{"msg": "OK"})
}
