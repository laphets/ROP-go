package user

import (
	. "git.zjuqsc.com/rop/ROP-go/handler"
	"git.zjuqsc.com/rop/ROP-go/model"
	"git.zjuqsc.com/rop/ROP-go/pkg/errno"
	"github.com/gin-gonic/gin"
)

func Info(c *gin.Context) {
	userId := c.GetInt("UserId")

	user, err := model.GetUserById(uint(userId))
	if err != nil {
		SendResponse(c, errno.ErrUserNotFound, nil)
		return
	}
	association, err := model.GetAssociationById(user.AssociationId)
	if err != nil {
		SendResponse(c, errno.ErrAssociationNotExist, nil)
		return
	}

	// Do not send password back
	user.Password = ""

	res := InfoResponse{
		UserModel: user,
		Association: association,
	}

	SendResponse(c, nil, res)
}