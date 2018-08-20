package form

import (
	"github.com/gin-gonic/gin"
	. "rop/handler"
	"rop/pkg/errno"
	"rop/model"
	json2 "encoding/json"
)

func Create(c *gin.Context) {
	req := CreateRequest{}
	if err := c.ShouldBindJSON(&req); err != nil {
		SendResponse(c, errno.ErrBind, err.Error())
		return
	}

	json, err := json2.Marshal(req)
	if err != nil {
		SendResponse(c, errno.ErrParam, nil)
		return
	}

	form := model.FormModel{
		Name:req.Name,
		Data:string(json),
	}

	if err := form.Create(); err != nil {
		SendResponse(c, errno.DBError, err.Error())
		return
	}
	SendResponse(c, nil, string(json))
}