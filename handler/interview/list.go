package interview

import (
	"github.com/gin-gonic/gin"
	"strconv"
	. "rop/handler"
	"rop/pkg/errno"
	"rop/model"
)

func List(c *gin.Context) {
	instanceId, err := strconv.ParseUint(c.Query("instanceId"), 10, 64)
	if err != nil {
		SendResponse(c, errno.ErrParam, err)
		return
	}
	interviewType, err := strconv.ParseUint(c.DefaultQuery("interview_type", "0"), 10, 64)
	if err != nil {
		SendResponse(c, errno.ErrParam, err)
		return
	}
	autoJoinable, err := strconv.ParseInt(c.DefaultQuery("auto_joinable", "0"), 10, 32)
	if err != nil {
		SendResponse(c, errno.ErrParam, err)
		return
	}
	department := c.DefaultQuery("department", "")

	if _, err := model.GetInstanceById(uint(instanceId)); err != nil {
		SendResponse(c, errno.ErrInstanceNotFound, nil)
		return
	}


	fulInterviews, err := model.ListFulInterview(uint(instanceId), &model.InterviewModel{
		InterviewType: uint(interviewType),
		Department: department,
		AutoJoinable: int(autoJoinable),
	})
	if err != nil {
		SendResponse(c, errno.DBError, err)
		return
	}

	SendResponse(c, nil, fulInterviews)
}