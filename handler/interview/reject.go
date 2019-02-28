package interview

import (
	"git.zjuqsc.com/rop/ROP-go/model"
	"git.zjuqsc.com/rop/ROP-go/pkg/errno"
	"git.zjuqsc.com/rop/ROP-go/service"
	"github.com/gin-gonic/gin"
	"strconv"
	. "git.zjuqsc.com/rop/ROP-go/handler"
)

func Reject(c *gin.Context) {
	intentId, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		SendResponse(c, errno.ErrParam, err)
		return
	}

	req := &RejectRequest{}
	if err := c.ShouldBindJSON(req); err != nil {
		SendResponse(c, errno.ErrBind, err.Error())
		return
	}

	uid := req.Uid

	if uid == "" {
		SendResponse(c, errno.ErrParam, "param is null")
		return
	}
	freshmanIdString, err := service.Decrypt(uid)
	if err != nil {
		SendResponse(c, errno.ErrParam, err)
		return
	}

	freshmanId, err := strconv.ParseUint(freshmanIdString, 10, 64)
	if err != nil {
		SendResponse(c, errno.ErrParam, err)
		return
	}

	_, err = model.GetFreshmanById(uint(freshmanId))
	if err != nil {
		SendResponse(c, errno.ErrUserNotFound, err.Error())
		return
	}

	intent, err := model.GetIntentByID(uint(intentId))
	if err != nil {
		SendResponse(c, errno.DBError, err.Error())
		return
	}

	intent.InterviewId = 0
	intent.TargetInterviewId = 0
	intent.SubStage = 1

	if err := intent.Update(); err != nil {
		SendResponse(c, errno.DBError, err.Error())
		return
	}

	SendResponse(c, nil, nil)
}