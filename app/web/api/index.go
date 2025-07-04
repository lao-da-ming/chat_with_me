package api

import "github.com/gin-gonic/gin"

type Index struct {
}

func NewIndex() *Index {
	return &Index{}
}
func (i *Index) Home(c *gin.Context) {
	c.JSON(200, gin.H{"message": "hello world"})
}
