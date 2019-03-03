package association

import (
	"fmt"
	. "git.zjuqsc.com/rop/ROP-go/handler"
	"git.zjuqsc.com/rop/ROP-go/model"
	"git.zjuqsc.com/rop/ROP-go/pkg/errno"
	"git.zjuqsc.com/rop/ROP-go/service"
	"github.com/gin-gonic/gin"
	"strconv"
)

func AddUser(c *gin.Context) {
	req := AddUserRequest{}
	if err := c.ShouldBindJSON(&req); err != nil {
		SendResponse(c, errno.ErrBind, err.Error())
		return
	}

	association, err := model.GetAssociationById(req.AssociationId)
	if err != nil {
		SendResponse(c, errno.ErrAssociationNotExist, err.Error())
		return
	}

	newuser := model.UserModel{
		AssociationId: req.AssociationId,
		ZJUid: req.ZJUid,
		Name: req.Name,
		Mobile: req.Mobile,
		Department: req.Department,
		AdminLevel: req.AdminLevel,
	}

	if err := newuser.Create(); err != nil {
		SendResponse(c, errno.DBError, err.Error())
		return
	}

	encryptedUserId, err := service.Encrypt(strconv.FormatUint(uint64(newuser.ID), 10))
	if err != nil {
		SendResponse(c, errno.ErrEncrypt, err)
		return
	}

	if _, err := service.SendRegisterNotice(newuser.Mobile, newuser.Name, c.GetString("UserName"), association.Name, fmt.Sprintf("https://rop.zjuqsc.com/console/register?uid=%s", encryptedUserId)); err != nil {
		SendResponse(c, errno.ErrSMS, err.Error())
		return
	}

	SendResponse(c, nil, encryptedUserId)
}

func GetUser(c *gin.Context) {
	adminLevel := c.GetInt("AdminLevel")

	if adminLevel >= 2 {
		// If system admin, return all users
		userList, err := model.ListUser()
		if err != nil {
			SendResponse(c, errno.DBError, err.Error())
			return
		}
		SendResponse(c, nil, userList)
	} else {
		userList, err := model.ListUserByAssociation(uint(c.GetInt("AssociationId")))
		if err != nil {
			SendResponse(c, errno.DBError, err.Error())
			return
		}
		SendResponse(c, nil, userList)
	}
}