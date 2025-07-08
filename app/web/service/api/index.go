package api

import (
	"chat_with_me/app/web/data"
	"chat_with_me/common/model/entity"
	"database/sql"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
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
	h.userRepo.Create(c, &entity.User{
		ID: time.Now().UnixMicro(),
		Name: sql.NullString{
			String: "劳达明",
			Valid:  true,
		},
		Attr: sql.NullString{},
		Path: sql.NullString{
			String: "ldm",
			Valid:  true,
		},
	})
	c.JSON(http.StatusOK, gin.H{"msg": "OK"})
}
