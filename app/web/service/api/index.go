package api

import (
	"chat_with_me/app/web/data"
	"chat_with_me/common/model/entity"
	"database/sql"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"strconv"
	"time"
)

type IndexController struct {
	logger   *zap.Logger
	userRepo *data.UserRepo
}

func NewIndexController(logger *zap.Logger, userRepo *data.UserRepo) *IndexController {
	return &IndexController{logger: logger, userRepo: userRepo}
}
func (h *IndexController) Home(c *gin.Context) {
	c.String(http.StatusOK, "Hello World!!")
}

func (h *IndexController) Create(c *gin.Context) {
	err := h.userRepo.Create(c, &entity.User{
		ID: time.Now().UnixMicro(),
		Name: sql.NullString{
			String: "劳达明",
			Valid:  true,
		},
		Attr: sql.NullString{
			String: "{}",
			Valid:  true,
		},
		Path: sql.NullString{
			String: "ldm",
			Valid:  true,
		},
	})
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "msg": "创建成功"})
}
func (h *IndexController) Update(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "msg": err.Error()})
		return
	}
	err = h.userRepo.Update(c, id, map[string]any{
		"name": sql.NullString{
			String: "哈哈",
			Valid:  true,
		},
	})
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "msg": "修改成功"})
}
func (h *IndexController) UpdateAttr(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "msg": err.Error()})
		return
	}
	err = h.userRepo.UpdateAttr(c, id, "attr",
		[]string{"profile", "info", "attr"},
		map[string]any{
			"name": "厉害",
			"age":  60,
		})
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "msg": "ok"})
}
