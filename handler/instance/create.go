package instance

import (
	"github.com/gin-gonic/gin"
	"rop/pkg/errno"
	"rop/model"
	. "rop/handler"
)

func Create(c *gin.Context) {

	req := CreateRequest{}
	if err := c.ShouldBindJSON(&req); err != nil {
		SendResponse(c, errno.ErrBind, err.Error())
		return
	}
	if !req.StartTime.Before(req.EndTime) {
		SendResponse(c, errno.ErrTime, nil)
		return
	}
	// TODO: Only Support >=1 intent now
	if req.MaxIntent <= 0 {
		SendResponse(c, errno.ErrParam, "IntentNum must > 0")
		return
	}
	//log.Debugf("%+v",req)
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
		MaxIntent:req.MaxIntent,
	}

	if err := ins.Create(); err != nil {
		SendResponse(c, errno.DBError, err.Error())
		return
	}

	SendResponse(c, nil, nil)
}