package form

import (
	"github.com/gin-gonic/gin"
	. "rop/handler"
	"rop/pkg/errno"
	"rop/model"
)

func Create(c *gin.Context) {
	req := CreateRequest{}
	if err := c.ShouldBindJSON(&req); err != nil {
		SendResponse(c, errno.ErrBind, err.Error())
		return
	}

	form := model.FormModel{
		Name:req.Name,
		Data:req.Data,
	}

	if err := form.Create(); err != nil {
		SendResponse(c, errno.DBError, err.Error())
		return
	}
	SendResponse(c, nil, nil)
}