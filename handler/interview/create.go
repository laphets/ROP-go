package interview

import (
	"github.com/gin-gonic/gin"
	. "git.zjuqsc.com/rop/ROP-go/handler"
	"git.zjuqsc.com/rop/ROP-go/pkg/errno"
	"git.zjuqsc.com/rop/ROP-go/model"
	"strconv"
)

func Create(c *gin.Context) {
	instanceId, err := strconv.ParseUint(c.Query("instanceId"), 10, 64)
	if err != nil {
		SendResponse(c, errno.ErrParam, err)
		return
	}
	if _, err := model.GetInstanceById(uint(instanceId)); err != nil {
		SendResponse(c, errno.ErrInstanceNotFound, nil)
		return
	}

	req := &CreateRequest{}
	if err := c.ShouldBindJSON(req); err != nil {
		SendResponse(c, errno.ErrBind, err.Error())
		return
	}

	if !req.StartTime.Before(req.EndTime) {
		SendResponse(c, errno.ErrTime, nil)
		return
	}

	if req.Capacity < 1 {
		SendResponse(c, errno.ErrParam, "Error capacity.")
		return
	}

	interview := &model.InterviewModel{
		InstanceId: uint(instanceId),
		Name:req.Name,
		InterviewType: req.InterviewType,
		Department:req.Department,
		Director:req.Director,
		AutoJoinable:req.AutoJoinable,
		Interviewers:req.Interviewers,
		StartTime:req.StartTime,
		EndTime:req.EndTime,
		Remark:req.Remark,
		Capacity: req.Capacity,
		Location: req.Location,
		NotAvailable: 0,
	}

	if err := interview.Create(); err != nil {
		SendResponse(c, errno.DBError, err.Error())
		return
	}
	SendResponse(c, nil, nil)
}