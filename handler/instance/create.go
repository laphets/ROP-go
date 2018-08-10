package instance

import (
	"github.com/gin-gonic/gin"
	. "rop/handler"
	"rop/pkg/errno"
	"rop/model"
	"github.com/lexkong/log"
)

func Create(c *gin.Context) {
	req := CreateRequest{}
	if err := c.ShouldBindJSON(&req); err != nil {
		SendResponse(c, errno.ErrBind, err.Error())
		return
	}
	log.Debugf("%+v",req)
	if _, err := model.GetInstanceByName(req.Name); err == nil {
		SendResponse(c, errno.DuplicateKey, nil)
		return
	}
	ins := model.InstanceModel{
		Name:req.Name,
		Remark:req.Remark,
		Association:req.Association,
		StartTime:req.StartTime,
		EndTime:req.EndTime,
		FormId:req.FormId,
	}

	if err := ins.Create(); err != nil {
		SendResponse(c, errno.DBError, err.Error())
		return
	}

	SendResponse(c, nil, nil)
}