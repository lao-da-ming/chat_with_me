package api

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type IndexController struct {
	logger *zap.Logger
}

func NewIndexController(logger *zap.Logger) *IndexController {
	return &IndexController{logger: logger}
}
func (i *IndexController) Home(c *gin.Context) {
	/*conn, ok := ws.ConnectPool.Load(c.Param("key"))
	if !ok {
		c.JSON(200, gin.H{"msg": "链接不存在"})
	}
	fd, ok := conn.(*ws.ConnObj)
	if !ok {
		c.JSON(200, gin.H{"msg": "断言链接失败"})
	}
	fd.WriteToSendChan([]byte("9999999999999999999999999"))*/
	i.logger.Error("hahahahah")
	c.JSON(200, gin.H{"msg": "OK"})
}
