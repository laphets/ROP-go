package interview

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"rop/pkg/errno"
	. "rop/handler"
	"rop/model"
	"rop/service"
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
	// TODO: Del this extra query
	interview, err := model.GetInterviewByID(uint(interviewId))
	if err != nil {
		SendResponse(c, errno.DBError, err.Error())
		return
	}


	fulInterview, err := model.GetFulInterviewByID(uint(interviewId))
	if err != nil {
		SendResponse(c, errno.DBError, err.Error())
		return
	}

	if len(fulInterview.Participants) + len(req.Intents) > fulInterview.Capacity {
		SendResponse(c, errno.ErrInterviewFull, nil)
		return
	}

	intent := &model.IntentModel{}
	for _, item := range req.Intents {

		fulIntent, err := model.GetFullIntentByID(item)
		if err != nil {
			SendResponse(c, errno.DBError, err.Error())
			return
		}

		if fulIntent.InstanceId != interview.InstanceId || fulIntent.Department != interview.Department || service.StateInNum(service.NextState(fulIntent.MainStage)) != interview.InterviewType {
			//log.Debugf("%s||%s", fulIntent.Department, interview.Department)
			SendResponse(c, errno.ErrOperation, "Instance or Department or stage Not Match")
			return
		}
		//log.Debugf("%+v", fulIntent.IntentModel)
		//log.Debugf("%d", item)
		intent.ID = item
		intent.InterviewId = uint(interviewId)
		intent.MainStage = service.NextState(fulIntent.MainStage)
		intent.SubStage = 1
		intent.TargetInterviewId = 0
		if err := intent.Update(); err != nil {
			SendResponse(c, errno.DBError, err.Error())
			return
		}
	}
	SendResponse(c, nil, nil)
}