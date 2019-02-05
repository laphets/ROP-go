package user

import (
	. "git.zjuqsc.com/rop/ROP-go/handler"
	"git.zjuqsc.com/rop/ROP-go/model"
	"git.zjuqsc.com/rop/ROP-go/pkg/errno"
	"github.com/gin-gonic/gin"
)

func Avatar(c *gin.Context) {
	req := AvatarRequest{}
	if err := c.ShouldBindJSON(&req); err != nil {
		SendResponse(c, errno.ErrBind, err.Error())
		return
	}
	ZJUid := c.GetString("ZJUid")
	if err := model.UpdateAvatar(ZJUid, req.URL); err != nil {
		SendResponse(c, errno.DBError, err.Error())
		return
	}
	SendResponse(c, nil, nil)
}