package ssr

import (
	"git.zjuqsc.com/rop/ROP-go/model"
	"git.zjuqsc.com/rop/ROP-go/pkg/errno"
	"git.zjuqsc.com/rop/ROP-go/service"
	"github.com/gin-gonic/gin"
	. "git.zjuqsc.com/rop/ROP-go/handler"
	"strconv"
)

func Register(c *gin.Context) {
	uid := c.DefaultQuery("uid", "")
	if uid == "" {
		SendResponse(c, errno.ErrParam, "param is null")
		return
	}
	userIdString, err := service.Decrypt(uid)
	if err != nil {
		SendResponse(c, errno.ErrParam, err)
		return
	}
	userId, err := strconv.ParseUint(userIdString, 10, 64)
	if err != nil {
		SendResponse(c, errno.ErrParam, err)
		return
	}

	user, err := model.GetUserById(uint(userId))

	if user.Password != "" {
		SendResponse(c, errno.ErrLinkInvaild, "The link has expired.")
		return
	}

	association, err := model.GetAssociationById(user.AssociationId)
	if err != nil {
		SendResponse(c, errno.ErrAssociationNotExist, err.Error())
		return
	}

	res := &RegisterResponse{
		Name: user.Name,
		Association: association,
	}

	SendResponse(c, nil, res)
}