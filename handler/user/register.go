package user

import (
	"git.zjuqsc.com/rop/ROP-go/model"
	"git.zjuqsc.com/rop/ROP-go/pkg/errno"
	"git.zjuqsc.com/rop/ROP-go/pkg/token"
	"git.zjuqsc.com/rop/ROP-go/service"
	"github.com/gin-gonic/gin"
	. "git.zjuqsc.com/rop/ROP-go/handler"
	"strconv"
)

func Register(c *gin.Context) {
	req := RegisterRequest{}
	if err := c.ShouldBindJSON(&req); err != nil {
		SendResponse(c, errno.ErrBind, err.Error())
		return
	}

	uid := req.Uid
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
	if err != nil {
		SendResponse(c, errno.ErrUserNotFound, err.Error())
		return
	}

	if user.Password != "" {
		SendResponse(c, errno.ErrLinkInvaild, "The user has registered")
		return
	}

	user.Password = service.SHAHash(req.Password)
	user.Department = req.Department
	if err := user.Update(); err != nil {
		SendResponse(c, errno.DBError, err.Error())
		return
	}

	JWT, err := token.Sign(token.Context{UserId:int(user.ID)}, "")
	if err != nil {
		SendResponse(c, errno.ErrToken, nil)
		return
	}

	SendResponse(c, nil, JWT)
}