package interview

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"rop/pkg/errno"
	. "rop/handler"
	"rop/model"
)

func Join(c *gin.Context) {
	interviewId, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		SendResponse(c, errno.ErrParam, err)
		return
	}

	req := &JoinRequest{}
	if err := c.ShouldBindJSON(req); err != nil {
		SendResponse(c, errno.ErrBind, err.Error())
		return
	}

	// Check for interview
	interview, err := model.GetInterviewByID(uint(interviewId));
	if err != nil {
		SendResponse(c, errno.DBError, err.Error())
		return
	}

	intent := &model.IntentModel{}
	for _, item := range req.Intents {
		if fulIntent, err := model.GetFullIntentByID(item); err != nil {
			SendResponse(c, errno.DBError, err.Error())
			return
		} else {
			if fulIntent.InstanceId != interview.InstanceId || fulIntent.Department != interview.Department {
				//log.Debugf("%s||%s", fulIntent.Department, interview.Department)
				SendResponse(c, errno.ErrOperation, "Instance or Department Not Match")
				return
			}
		}

		//log.Debugf("%d", item)
		intent.ID = item
		intent.InterviewId = uint(interviewId)
		if err := intent.Update(); err != nil {
			SendResponse(c, errno.DBError, err.Error())
			return
		}
	}
	SendResponse(c, nil, nil)
}