package user

import (
	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
	. "rop/handler"
	"rop/model"
	"rop/pkg/errno"
)

func Info(c *gin.Context) {
	ZJUid := c.GetString("ZJUid")
	log.Debugf(ZJUid)
	user, err := model.GetUserByZJUid(ZJUid)
	if err != nil {
		SendResponse(c, errno.ErrUserNotFound, nil)
		return
	}
	SendResponse(c, nil, user)
}