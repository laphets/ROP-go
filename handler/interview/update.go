package interview

import (
	"strconv"
	"git.zjuqsc.com/rop/ROP-go/pkg/errno"
	"git.zjuqsc.com/rop/ROP-go/model"
	. "git.zjuqsc.com/rop/ROP-go/handler"
	"github.com/gin-gonic/gin"
)

func Update(c *gin.Context) {
	interviewId, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		SendResponse(c, errno.ErrParam, err)
		return
	}

	interview := model.InterviewModel{}
	if err := c.ShouldBindJSON(&interview); err != nil {
		SendResponse(c, errno.ErrBind, err)
		return
	}

	// Check for not null
	if interview.Department != "" || interview.InstanceId != 0 {
		SendResponse(c, errno.ErrOperation, nil)
		return
	}


	interview.ID = uint(interviewId)

	if err := interview.Update(); err != nil {
		SendResponse(c, errno.DBError, err.Error())
		return
	}
	SendResponse(c, nil, nil)
}
