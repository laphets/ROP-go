package intent

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"rop/pkg/errno"
	. "rop/handler"
	"rop/model"
)

func Reject(c *gin.Context) {
	intentId, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		SendResponse(c, errno.ErrParam, err)
		return
	}

	intent, err := model.GetIntentByID(uint(intentId))
	if err != nil {
		SendResponse(c, errno.DBError, err)
		return
	}

	if intent.SubStage != 2 {
		SendResponse(c, errno.ErrOperation, "Not in target status.")
		return
	}

	intent.SubStage = 1
	intent.TargetInterviewId = 0
	if err := intent.Update(); err != nil {
		SendResponse(c, errno.DBError, err)
		return
	}

	SendResponse(c, nil, nil)
}