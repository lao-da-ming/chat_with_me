package api

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

type IndexController struct {
	logger *zap.Logger
}

func NewIndexController(logger *zap.Logger) *IndexController {
	return &IndexController{logger: logger}
}
func (i *IndexController) Home(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"msg": "OK"})
}
