package api

import "github.com/gin-gonic/gin"

type IndexController struct {
}

func NewIndexController() *IndexController {
	return &IndexController{}
}
func (i *IndexController) Home(c *gin.Context) {
	c.JSON(200, gin.H{"message": "hello world"})
}
