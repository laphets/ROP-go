package instance

import (
	. "git.zjuqsc.com/rop/ROP-go/handler"
	"git.zjuqsc.com/rop/ROP-go/model"
	"git.zjuqsc.com/rop/ROP-go/pkg/errno"
	"github.com/gin-gonic/gin"
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

	if _, err := model.GetFormByID(req.FormId); err != nil {
		SendResponse(c, errno.ErrFormNotFound, nil)
		return
	}

	association, err := model.GetAssociationById(uint(c.GetInt("AssociationId")))
	if err != nil {
		SendResponse(c, errno.ErrAssociationNotExist, err.Error())
		return
	}

	ins := model.InstanceModel{
		Name:        req.Name,
		AssociationId: association.ID,
		Remark:      req.Remark,
		Association: association.Name,
		StartTime:   req.StartTime,
		EndTime:     req.EndTime,
		FormId:      req.FormId,
		MaxIntent:   req.MaxIntent,
	}

	if err := ins.Create(); err != nil {
		SendResponse(c, errno.DBError, err.Error())
		return
	}

	SendResponse(c, nil, nil)
}
