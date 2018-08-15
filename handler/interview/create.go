package interview

import (
	"github.com/gin-gonic/gin"
	. "rop/handler"
	"rop/pkg/errno"
	"rop/model"
)

func Submit(c *gin.Context) {
	req := &CreateRequest{}
	if err := c.ShouldBindJSON(req); err != nil {
		SendResponse(c, errno.ErrBind, err.Error())
		return
	}

	if !req.StartTime.Before(req.EndTime) {
		SendResponse(c, errno.ErrTime, nil)
		return
	}

	instanceId := req.InstanceId

	if _, err := model.GetInstanceById(instanceId); err != nil {
		SendResponse(c, errno.ErrInstanceNotFound, nil)
		return
	}

	interview := &model.InterviewModel{
		InstanceId: req.InstanceId,
		Name:req.Name,
		InterviewType: req.InterviewType,
		Department:req.Department,
		Director:req.Director,
		AutoJoinable:req.AutoJoinable,
		Interviewers:req.Interviewers,
		StartTime:req.StartTime,
		EndTime:req.EndTime,
		Remark:req.Remark,
	}

	if err := interview.Create(); err != nil {
		SendResponse(c, errno.DBError, nil)
		return
	}
	SendResponse(c, nil, nil)
}